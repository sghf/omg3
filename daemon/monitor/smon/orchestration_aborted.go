package smon

func (o *smon) orchestrateAborted() {
	if !o.isConvergedGlobalExpect() {
		return
	}
	o.log.Info().Msg("abort orchestration: unset global expect")
	o.change = true
	o.state.GlobalExpect = globalExpectUnset
	o.updateIfChange()
}

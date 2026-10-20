package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = -4
    var v3 int = 4294967295
    var v4 uint8 = 255
    var v5 uint16 = 32767
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(fi6)
        var v7 uint8 = v4
        var v8 int32 = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: (v3 ^ (v3 + (v3 / ((v3) | 1)))), F1: true}
    v4 += 166
    var v10 uint8 = (v4 & (v4 & (v4 ^ (v4 | v4))))
    fmt.Println(v4, v5)
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 int16 = -32768
        v3 = 356
        var v13 uint32 = 42
        _ = fi11
        _ = v12
        _ = v13
    }
    if false {
        for fi14 := 0; fi14 < 5; fi14++ {
            if (("hello" + ("\n\t" + (("a" + "") + "a"))) == ("" + "\n\t")) {
                if false {
                    fmt.Println(v3, v3)
                    v15 := GenS2{F0: -2147483648, F1: 1000}
                    fmt.Println(v10, v10, v10)
                    _ = v15
                } else {
                    switch fi14 {
                    case 1:
                        for fi16 := 0; fi16 < 4; fi16++ {
                            v17 := GenS2{F0: -1, F1: 1000}
                            v2 = 87
                            fmt.Println(v3, v2, fi16)
                            _ = fi16
                            _ = v17
                        }
                    case 32767:
                        for fi18 := 0; fi18 < 11; fi18++ {
                            v3 -= (-(fi18))
                            if false {
                                _ = len(v1)
                            } else {
                                fmt.Println(v10, v5, v3)
                                v19 := []int{(fi18 + -807), v3, 4294967295, (-(v3))}
                                fmt.Println(fi14, v10)
                                _ = v19
                            }
                            fmt.Println(v10)
                            _ = fi18
                        }
                    default:
                        var v20 int64 = -1
                        _ = v20
                    }
                }
                if false {
                    var v21 int8 = -63
                    var v22 int32 = 236
                    v2 -= v2
                    _ = v21
                    _ = v22
                } else {
                    v5 = 358
                }
                v23 := GenS1{F0: (fi14 + fi14), F1: (v5 > (v5 >> ((0) & 63)))}
                _ = v23
            }
            _ = fi14
        }
    } else {
        fmt.Println(v2, v3, v4)
        if false {
            var v24 bool = (((!true) && false) && false)
            var v25 uint32 = 127
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 int8 = v2
                _ = fi26
                _ = v27
            }
            _ = v24
            _ = v25
        } else {
            var v28 float32 = (1.5 + -2.25)
            v1[("" + "foo bar")] = v3
            if false {
                v5 -= v5
                var v29 bool = (false || false)
                _ = v29
            }
            _ = v28
        }
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        if true {
            v10 = (v10 & v4)
            v2 += -128
            if true {
                var v31 int8 = (v2 >> ((1) & 63))
                _ = v31
            } else {
                fmt.Println(fi30, v2, v10)
            }
        } else {
            fmt.Println(v4, v5)
            v32 := GenS3{F0: 481, F1: 45, F2: 2147483648}
            _ = v32
        }
        v10 -= (v10 << ((1) & 63))
        if (("foo bar" + ("a" + "\n\t")) >= "hello") {
            v33 := GenS2{F0: 0, F1: 133}
            if (((v2 + 1) == v2) || false) {
                v34 := GenS3{F0: -601, F1: 2, F2: 967}
                v35 := GenS3{F0: 2, F1: 42, F2: 405}
                v5 = (v5 * 32767)
                _ = v34
                _ = v35
            } else {
                if true {
                    v36 := []int{(v3 + (v3 + (fi30 + v3)))}
                    fmt.Println(v2, v4, v10)
                    for fi37 := 0; fi37 < 1; fi37++ {
                        v4 = 42
                        _ = fi37
                    }
                    _ = v36
                }
                if false {
                    v38 := GenS1{F0: (v3 % ((fi30) | 1)), F1: true}
                    _ = v38
                } else {
                    v4 += 2
                }
                var v39 uint = 32
                _ = v39
            }
            v40 := GenS2{F0: 0, F1: 32767}
            _ = v33
            _ = v40
        }
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}

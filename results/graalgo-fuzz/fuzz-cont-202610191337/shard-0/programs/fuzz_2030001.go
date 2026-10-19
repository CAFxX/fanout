package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int; F1 int8 }
func genFn0(p0 uint8) int16 {
    var v5 float32 = float32(float64((-2.25 + ((1.5 * (-2.25 / 1.5)) + 1e10))))
    v6 := GenS1{F0: 2147483648, F1: p0}
    _ = p0
    _ = v5
    _ = v6
    return -32768
}

func main() {
    var v1 bool = false
    v2 := []uint32{1, 0}
    v3 := []int{}
    var v4 uint16 = 0
    if v1 {
        v7 := GenS1{F0: 0, F1: 128}
        var v8 int64 = 42
        _ = v7
        _ = v8
    } else {
        fmt.Println(v4, v1)
        var v9 int = -989
        _ = v9
    }
    if len(v2) > 0 {
        v2[0] = 526
    }
    switch ("foo bar" + "") {
    case "hello":
        if v1 {
            for fi10 := 0; fi10 < 12; fi10++ {
                fmt.Println(v4)
                for fi11 := 0; fi11 < 7; fi11++ {
                    v12 := GenS2{F0: (fi10 << ((0) & 63)), F1: 127}
                    var v13 float32 = 1e10
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi10
            }
            for fi14 := 0; fi14 < 19; fi14++ {
                switch fi14 {
                case 2147483647:
                    v4 -= 255
                default:
                    v1 = ((((2 != 313) && (v1 && ((1e10 == float32(0.1)) && (42 >= 834)))) || v1) && (("foo bar" < ("\n\t" + ("hello" + "a"))) && v1))
                }
                switch "a" {
                case "":
                    fmt.Println(v4)
                case "a":
                    var v15 float32 = 1.5
                    _ = v15
                default:
                    var v16 float32 = (0.1 / 1.0)
                    _ = v16
                }
                if v1 {
                    fmt.Println(v4)
                }
                _ = fi14
            }
        }
    case "\n\t":
        fmt.Println(v4, v1)
    case "a":
        fmt.Println(v4)
    }
    v1 = v1
    if len(v3) > 0 {
        v3[0] = 32767
    }
    var v17 uint16 = (v4 / ((0) | 1))
    var v18 int16 = 42
    for fi19 := 0; fi19 < 18; fi19++ {
        if (fi19 != fi19) {
            var v20 bool = (!(42 >= 2147483647))
            v17 = 255
            _ = v20
        } else {
            if false {
                for fi21 := 0; fi21 < 20; fi21++ {
                    var v22 uint16 = v4
                    if len(v3) > 0 {
                        v3[0] = (fi19 >> ((42) & 63))
                    }
                    _ = fi21
                    _ = v22
                }
                fmt.Println(v4, v4, v17)
                _ = len(v3)
            }
            var v23 string = "a"
            _ = v23
        }
        for fi24 := 0; fi24 < 17; fi24++ {
            v1 = v1
            if (!v1) {
                var v25 uint16 = v17
                if true {
                    var v26 int64 = 9223372036854775807
                    switch (fi24 % (((-(fi19))) | 1)) {
                    case 64:
                        _ = len(v3)
                    case 4294967295:
                        var v27 uint32 = 4294967295
                        var v28 uint32 = (v27 - (v27 & (v27 % ((v27) | 1))))
                        var v29 int16 = -1
                        _ = v27
                        _ = v28
                        _ = v29
                    case 1:
                        if true {
                            if false {
                                var v30 int64 = 1
                                _ = len(v3)
                                var v31 int64 = -9223372036854775808
                                _ = v30
                                _ = v31
                            }
                            if false {
                                var v32 int = (-(fi24))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v1 = v1
                                _ = v32
                            } else {
                                if (("" + "hello") >= "foo bar") {
                                    v1 = v1
                                    switch int(v25) {
                                    case -2147483648:
                                        if len(v3) > 0 {
                                            v3[0] = fi19
                                        }
                                    }
                                    var v33 int16 = (^(v18))
                                    _ = v33
                                }
                                var v34 bool = v1
                                var v35 uint32 = 11
                                _ = v34
                                _ = v35
                            }
                        }
                        if false {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                    v1 = false
                    _ = v26
                }
                var v36 float64 = (1.5 / -2.25)
                fmt.Println(fi19, fi19)
                switch ("foo bar" + ("foo bar" + "")) {
                case "foo bar":
                    var v37 float32 = 1e10
                    _ = v37
                case "":
                    fmt.Println(v18, v1)
                }
                _ = v25
                _ = v36
            } else {
                v4 -= 1
                v17 += (v17 % ((v4) | 1))
                var v38 float32 = (((0.1 / 1.5) * 0.1) + 0.0)
                _ = v38
            }
            _ = fi24
        }
        _ = fi19
    }
    var v39 int = -841
    var v40 uint64 = 2
    if len(v3) > 0 {
        v3[0] = v39
    }
    for fi41 := 0; fi41 < 20; fi41++ {
        for fi42 := 0; fi42 < 3; fi42++ {
            fmt.Println(v39, v4)
            _ = len(v2)
            v3 = append(v3, (v39 * int(v18)))
            _ = fi42
        }
        v2 = append(v2, 685)
        _ = fi41
    }
    v40 -= 605
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
}

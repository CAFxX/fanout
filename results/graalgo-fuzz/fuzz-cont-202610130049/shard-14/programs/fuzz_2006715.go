package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int; F1 uint64; F2 int8 }
func genFn0(p0 uint, p1 uint) int32 {
    switch 767 {
    case 32767:
        fmt.Println(p1, p1)
    case 128:
        var v4 int = 106
        _ = v4
    case 4294967295:
        if false {
            for fi5 := 0; fi5 < 20; fi5++ {
                if true {
                    var v6 uint = p0
                    _ = v6
                }
                _ = fi5
            }
            p0 = 1
            p1 += 65535
        }
    default:
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := GenS1{F0: (fi7 + 4294967295), F1: true}
            var v9 string = "\n\t"
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v10 bool = true
    _ = p0
    _ = p1
    _ = v10
    return -49
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    fmt.Println(v1)
    v1 = v1
    v1 += (v1 - (v1 >> ((0) & 63)))
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            v1 -= (v1 & v1)
            var v13 int8 = 127
            if (v13 == 100) {
                switch fi12 {
                case -819:
                    if (255 <= 18446744073709551615) {
                        if ((-2147483648 > 42) || true) {
                            if false {
                                v1 = (^(v1))
                            }
                            if true {
                                var v14 uint32 = 312
                                _ = len(v2)
                                _ = v14
                            }
                            var v15 uint16 = 390
                            _ = v15
                        } else {
                            v1 -= -9223372036854775808
                            v1 = v1
                        }
                        var v16 uint = 18446744073709551615
                        var v17 int64 = (v1 >> ((v16) & 63))
                        _ = v16
                        _ = v17
                    } else {
                        v18 := []int{4294967295}
                        if (((!false) || (721 < 345)) && ((-202 < 0) || (65535 != 1))) {
                            var v19 string = "\n\t"
                            var v20 uint64 = 18446744073709551615
                            v21 := GenS2{F0: 4294967295}
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            var v22 float64 = (-2.25 - (float64(((1e10 + 1.5) * 0.1)) * -2.25))
                            if false {
                                v22 -= (v22 / -2.25)
                                var v23 uint16 = 255
                                _ = v23
                            } else {
                                v24 := GenS3{F0: int32(v13), F1: ""}
                                _ = v24
                            }
                            _ = v22
                        }
                        var v25 uint8 = 1
                        _ = v18
                        _ = v25
                    }
                case -2147483648:
                    switch "foo bar" {
                    case "\n\t":
                        v26 := []int{32768, (fi11 % ((fi12) | 1)), fi11}
                        _ = v26
                    case "a":
                        var v27 uint64 = 2
                        _ = v27
                    case "foo bar":
                        var v28 uint8 = 255
                        _ = v28
                    default:
                        v13 = int8(v1)
                    }
                }
            } else {
                v29 := GenS1{F0: (fi12 - (-(fi12))), F1: (4294967295 != 948)}
                _ = v29
            }
            v1 += (-(v1))
            _ = fi12
            _ = v13
        }
        switch 4294967295 {
        case -1:
            v1 = (v1 + (v1 << ((176) & 63)))
        case -919:
            for fi30 := 0; fi30 < 11; fi30++ {
                if (!false) {
                    var v31 int8 = 100
                    var v32 int64 = v1
                    _ = v31
                    _ = v32
                } else {
                    v1 += (v1 & 837)
                }
                _ = fi30
            }
        case 121:
            v33 := GenS4{F0: fi11, F1: 2, F2: 0}
            _ = v33
        }
        v34 := GenS4{F0: -34, F1: 739, F2: 0}
        _ = fi11
        _ = v34
    }
    if ("hello" <= (("a" + "a") + "")) {
        var v35 float64 = 0.1
        _ = v35
    } else {
        v36 := []int{0, -1}
        _ = v36
    }
    if (42 <= genFn0(0, 18446744073709551615)) {
        var v37 string = ("\n\t" + "")
        if (0 == -2147483648) {
            v37 = (((("foo bar" + ((("foo bar" + v37) + "\n\t") + "\n\t")) + (v37 + v37)) + ("hello" + "foo bar")) + (v37 + "\n\t"))
            var v38 int32 = genFn0(65535, 2)
            _ = v38
        }
        _ = v37
    } else {
        fmt.Println(v1)
    }
    var v39 float64 = 0.0
    fmt.Println(v1, v1)
    var v40 uint = 1
    for fi41 := 0; fi41 < 5; fi41++ {
        for fi42 := 0; fi42 < 9; fi42++ {
            var v43 bool = false
            var v44 int = 1
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi41
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
}

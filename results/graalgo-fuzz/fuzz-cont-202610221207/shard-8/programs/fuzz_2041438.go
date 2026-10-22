package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int16; F1 string }

func main() {
    v1 := []string{("a" + "hello")}
    var v2 int64 = -1
    v3 := []string{""}
    var v4 int32 = -820
    var v5 float32 = float32(0.0)
    _ = len(v3)
    var v6 int64 = v2
    for fi7 := 0; fi7 < 8; fi7++ {
        v3 = append(v3, ("\n\t" + (("hello" + "a") + "\n\t")))
        v4 += (v4 << ((2) & 63))
        _ = fi7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        v4 = -746
        _ = fi8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v2 = v2
        _ = fi9
    }
    v10 := GenS1{F0: (v2 * (v6 << ((4294967295) & 63)))}
    v4 = v4
    if false {
        for fi11 := 0; fi11 < 4; fi11++ {
            if (856 == 770) {
                if (!(!true)) {
                    v3 = append(v3, "a")
                    v1 = append(v1, "")
                    v3 = append(v3, ("" + ("a" + "hello")))
                } else {
                    v4 += (v4 % ((42) | 1))
                    var v12 bool = (1 == 42)
                    _ = v12
                }
                var v13 uint64 = 343
                var v14 int64 = (v2 * (v6 - v6))
                _ = v13
                _ = v14
            }
            switch 115 {
            case 69:
                v6 = (v2 * 86)
            case -2147483648:
                for fi15 := 0; fi15 < 13; fi15++ {
                    v5 -= 0.1
                    _ = fi15
                }
            default:
                if false {
                    _ = len(v1)
                    for fi16 := 0; fi16 < 15; fi16++ {
                        _ = len(v3)
                        v5 -= v5
                        _ = fi16
                    }
                    for fi17 := 0; fi17 < 17; fi17++ {
                        v1 = append(v1, "\n\t")
                        v4 += v4
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                        _ = fi17
                    }
                }
            }
            v18 := GenS2{F0: 1, F1: ("hello" + "a")}
            _ = fi11
            _ = v18
        }
        if true {
            var v19 uint = 4294967295
            _ = v19
        } else {
            if (!false) {
                if ((float32((3.14159 - (0.0 * ((3.14159 / float64((float32(3.14159) * v5))) * 1e10)))) / v5) > ((((v5 * float32(1e10)) * v5) - v5) * v5)) {
                    var v20 uint8 = 42
                    v20 = v20
                    _ = v20
                } else {
                    v5 += v5
                    v4 = (^(v4))
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    if true {
        v21 := GenS2{F0: v2, F1: ""}
        if false {
            v22 := GenS2{F0: v2, F1: ("" + "")}
            for fi23 := 0; fi23 < 13; fi23++ {
                v4 = (v4 * v4)
                v4 += (^(v4))
                _ = fi23
            }
            var v24 int32 = v4
            _ = v22
            _ = v24
        } else {
            v25 := GenS3{F0: 865, F1: 1}
            v6 = v6
            if len(v3) > 0 {
                v3[0] = ("hello" + "a")
            }
            _ = v25
        }
        if false {
            var v26 int32 = v4
            _ = v26
        } else {
            if len(v3) > 0 {
                v3[0] = (("hello" + "hello") + "\n\t")
            }
            v27 := GenS2{F0: v2, F1: "hello"}
            _ = v27
        }
        _ = v21
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        var v29 float32 = (v5 + ((v5 - (-2.25 + ((1e10 + v5) - (((v5 * ((1.5 + 1.5) - (1.5 + v5))) - v5) / v5)))) * v5))
        v30 := GenS4{F0: -406, F1: ("" + "")}
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 float32 = ((((((float32(-2.25) + v5) / v5) - float32((float64(v5) + (((3.14159 * 1.5) * float64(v5)) - 1.5)))) * -2.25) - -2.25) - (v5 * ((v5 + 1e10) + (1e10 + v5))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}

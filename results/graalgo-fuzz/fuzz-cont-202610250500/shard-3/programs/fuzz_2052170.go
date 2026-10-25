package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int64 }
func genFn0() string {
    var v4 float32 = float32(0.0)
    for fi5 := 0; fi5 < 19; fi5++ {
        v4 = 3.14159
        var v6 uint16 = 32767
        v7 := []int{}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint16 = 0
    if false {
        v8 -= (v8 << ((571) & 63))
        if false {
            fmt.Println(v8)
        }
        v9 := []int{943}
        _ = v9
    } else {
        v8 += 0
    }
    _ = v4
    _ = v8
    return "hello"
}
func genFn1() string {
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int32 = 42
        v11 -= (v11 / (((-(v11))) | 1))
        _ = fi10
        _ = v11
    }
    v12 := []int{-984}
    v12 = append(v12, 9223372036854775807)
    v13 := GenS1{F0: -509}
    _ = v12
    _ = v13
    return "hello"
}

func main() {
    var v1 bool = false
    var v2 uint = 639
    var v3 int64 = 1
    v14 := GenS1{F0: -796}
    v15 := []int{-1, 1}
    if true {
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := GenS1{F0: int32(v3)}
            var v18 uint16 = 0
            v18 = v18
            _ = fi16
            _ = v17
            _ = v18
        }
    } else {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        if (!v1) {
            if (-763 > 2147483647) {
                for fi19 := 0; fi19 < 6; fi19++ {
                    v3 = (v3 % ((1) | 1))
                    _ = fi19
                }
                v1 = v1
            } else {
                for fi20 := 0; fi20 < 17; fi20++ {
                    v15 = append(v15, (-(fi20)))
                    _ = fi20
                }
                for fi21 := 0; fi21 < 6; fi21++ {
                    v1 = (false && v1)
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    _ = fi21
                }
            }
        }
        v22 := GenS1{F0: 0}
        _ = v22
    }
    v23 := GenS1{F0: 393}
    if (0.0 != 3.14159) {
        if (1 < 128) {
            for fi24 := 0; fi24 < 17; fi24++ {
                if v1 {
                    v25 := GenS1{F0: 334}
                    if v1 {
                        switch (fi24 - fi24) {
                        case 9223372036854775807:
                            fmt.Println(v1)
                        case 558:
                            v26 := GenS1{F0: 0}
                            _ = v26
                        default:
                            v2 -= (v2 - (v2 - v2))
                        }
                    }
                    _ = v25
                } else {
                    v2 -= v2
                    var v27 int16 = -781
                    _ = v27
                }
                fmt.Println(fi24)
                _ = fi24
            }
        }
        switch ("" + "\n\t") {
        case "a":
            switch -644 {
            case 0:
                v28 := GenS2{F0: 104, F1: 2}
                _ = v28
            }
        }
        var v29 int64 = (v3 & v3)
        _ = v29
    }
    v3 -= -1
    v30 := GenS1{F0: 808}
    if false {
        if v1 {
            for fi31 := 0; fi31 < 5; fi31++ {
                if (42 > 18446744073709551615) {
                    v32 := GenS3{F0: 42}
                    _ = v32
                } else {
                    if len(v15) > 0 {
                        v15[0] = fi31
                    }
                }
                _ = fi31
            }
            for fi33 := 0; fi33 < 8; fi33++ {
                v3 -= (-(v3))
                for fi34 := 0; fi34 < 16; fi34++ {
                    v35 := GenS3{F0: -372}
                    var v36 uint8 = 128
                    v37 := GenS2{F0: (v36 % ((v36) | 1)), F1: 32767}
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
                v2 -= (^(v2))
                _ = fi33
            }
        }
        v2 = (v2 + v2)
    }
    var v38 uint64 = 42
    if len(v15) > 0 {
        v15[0] = 2147483647
    }
    v15 = append(v15, 1)
    for fi39 := 0; fi39 < 9; fi39++ {
        var v40 uint64 = v38
        v1 = (!true)
        _ = fi39
        _ = v40
    }
    if len(v15) > 0 {
        v15[0] = 2147483647
    }
    v38 -= 2
    if ((!false) || v1) {
        v41 := GenS1{F0: 42}
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v38
}

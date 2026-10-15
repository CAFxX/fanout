package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 string }
type GenS2 struct { F0 uint32; F1 string; F2 uint32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint64) uint32 {
    p0 = (p0 & p0)
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    v1 := []uint32{2, 490, 2}
    v2 := []int{9223372036854775807, -377}
    var v3 float64 = 3.14159
    var v4 uint = 255
    var v5 uint8 = 145
    v6 := []string{("\n\t" + ("hello" + ("foo bar" + "a"))), (((((("foo bar" + "") + ("a" + "\n\t")) + "") + "\n\t") + "\n\t") + "hello"), "foo bar", "foo bar"}
    var v7 int32 = -94
    if len(v6) > 0 {
        v6[0] = "a"
    }
    fmt.Println(v5, v4)
    if len(v1) > 0 {
        v1[0] = 2
    }
    v3 = v3
    for fi8 := 0; fi8 < 6; fi8++ {
        if ((32767 == 32767) || (!(((v4 * v4) < v4) && ((17 < (^(v5))) && (-9223372036854775808 < -980))))) {
            v3 -= (v3 / v3)
            v3 -= v3
            if false {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        }
        for fi9 := 0; fi9 < 1; fi9++ {
            v5 -= (v5 + v5)
            _ = fi9
        }
        _ = fi8
    }
    var v10 bool = true
    if false {
        var v11 uint8 = v5
        if v10 {
            v12 := GenS1{F0: 0, F1: (^(v11)), F2: "a"}
            if (v10 || true) {
                for fi13 := 0; fi13 < 9; fi13++ {
                    fmt.Println(v4, v7, fi13)
                    _ = fi13
                }
                v7 += v7
            } else {
                if false {
                    if v10 {
                        var v14 int = int(v4)
                        v5 = 1
                        _ = v14
                    } else {
                        v2 = append(v2, 2147483647)
                    }
                    switch -615 {
                    case 516:
                        if v10 {
                            v15 := GenS1{F0: 42, F1: v11, F2: "\n\t"}
                            fmt.Println(v7)
                            _ = v15
                        }
                    default:
                        v16 := GenS2{F0: 42, F1: ("foo bar" + ("foo bar" + "")), F2: 4294967295}
                        _ = v16
                    }
                    v5 -= 241
                }
            }
            fmt.Println(v5, v5)
            _ = v12
        } else {
            var v17 float64 = v3
            _ = v17
        }
        _ = v11
    }
    if (214 > (^(v5))) {
        var v18 int8 = 1
        if (!(!false)) {
            var v19 uint64 = 435
            if len(v1) > 0 {
                v1[0] = 65535
            }
            v10 = (!(596 >= 2))
            _ = v19
        }
        _ = v18
    } else {
        var v20 uint8 = v5
        if (!((v10 || v10) && true)) {
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
            v20 = (v5 >> ((65535) & 63))
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            v4 = 476
            var v22 int64 = -9223372036854775808
            var v23 uint = 65535
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = v20
    }
    v7 -= -2147483648
    for fi24 := 0; fi24 < 9; fi24++ {
        switch -87 {
        case 233:
            v25 := GenS3{F0: 0}
            _ = v25
        case 42:
            var v26 float64 = 0.0
            _ = v26
        }
        _ = fi24
    }
    v10 = v10
    v27 := GenS2{F0: genFn0(373), F1: "a", F2: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

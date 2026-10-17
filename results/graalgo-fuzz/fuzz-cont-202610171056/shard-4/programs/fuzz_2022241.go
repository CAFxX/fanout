package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 int16; F2 string }
type GenS5 struct { F0 uint; F1 int8 }
func genFn0(p0 uint) uint8 {
    v5 := GenS1{F0: p0, F1: true, F2: -32768}
    p0 = (p0 + (p0 >> (((p0 | p0)) & 63)))
    for fi6 := 0; fi6 < 2; fi6++ {
        if ((-2.25 - (1.5 - float64(-2.25))) == (0.0 + 0.0)) {
            switch "hello" {
            case "foo bar":
                v7 := GenS2{F0: 42}
                _ = v7
            case "hello":
                var v8 int = fi6
                _ = v8
            default:
                v9 := []int{(-(fi6)), fi6, -575, (-(fi6))}
                _ = v9
            }
        }
        _ = fi6
    }
    _ = p0
    _ = v5
    return 134
}

func main() {
    v1 := []uint32{42, 65535}
    var v2 int = -9223372036854775808
    var v3 bool = (!(1 > -74))
    v4 := []uint32{327, 2147483647}
    var v10 int64 = -690
    v11 := GenS1{F0: 18446744073709551615, F1: (!(false || v3)), F2: -547}
    for fi12 := 0; fi12 < 5; fi12++ {
        if true {
            for fi13 := 0; fi13 < 10; fi13++ {
                v10 = v10
                var v14 float64 = ((1.5 - (0.0 - (1e10 / float64((0.1 - 1.5))))) - (1.5 - (1.5 - 1.5)))
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 8; fi15++ {
                var v16 float64 = 1e10
                if (!((0.1 + 1e10) == 0.1)) {
                    if (2 != 2147483647) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                } else {
                    v2 = 9223372036854775807
                }
                v17 := GenS1{F0: 255, F1: v3, F2: -1}
                v3 = v3
                _ = fi15
                _ = v16
                _ = v17
            }
        } else {
            _ = len(v4)
            v10 = -9223372036854775808
            if v3 {
                if v3 {
                    fmt.Println(v3, v2, v2)
                    var v18 uint64 = 2
                    var v19 int8 = -1
                    _ = v18
                    _ = v19
                }
            }
        }
        v10 = v10
        v20 := GenS2{F0: 1}
        _ = fi12
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 uint32 = 905
        _ = fi21
        _ = v22
    }
    v23 := GenS3{F0: v10}
    v24 := GenS4{F0: (v10 % (((-(v10))) | 1)), F1: 837, F2: "hello"}
    if ((4294967295 <= 4294967295) && v3) {
        _ = len(v1)
    } else {
        var v25 string = ""
        for fi26 := 0; fi26 < 13; fi26++ {
            var v27 int32 = 0
            for fi28 := 0; fi28 < 8; fi28++ {
                var v29 int = (fi26 >> ((107) & 63))
                var v30 int64 = (v10 << ((42) & 63))
                var v31 int8 = 0
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            if len(v4) > 0 {
                v4[0] = 65535
            }
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        if (!false) {
            if (false && (255 >= 42)) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi32 := 0; fi32 < 5; fi32++ {
                    v4 = append(v4, 285)
                    _ = fi32
                }
                fmt.Println(v3, v3)
            }
        }
    case "a":
        v33 := GenS5{F0: 4294967295, F1: -88}
        _ = v33
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        _ = len(v4)
        _ = fi34
    }
    v3 = v3
    v3 = (128 > 92)
    v2 += v2
    v3 = ((65535 <= 448) && (775 == -2147483648))
    v1 = append(v1, 0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}

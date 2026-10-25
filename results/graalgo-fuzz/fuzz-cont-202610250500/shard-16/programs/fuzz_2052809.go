package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int32; F2 int16 }
func genFn0(p0 bool) int64 {
    p0 = ((-2.25 * float64(-2.25)) == 0.1)
    var v5 float32 = (0.0 - -2.25)
    _ = p0
    _ = v5
    return 619
}
func genFn1(p0 int, p1 int64, p2 uint) string {
    fmt.Println(p1)
    p1 = (p1 << ((p2) & 63))
    p0 -= (p0 << (((p2 | p2)) & 63))
    var v6 uint16 = 65535
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"hello", "hello"}
    var v3 string = ("\n\t" + "")
    var v4 int32 = 972
    for fi8 := 0; fi8 < 17; fi8++ {
        v2 = append(v2, (((v3 + v3) + v3) + "foo bar"))
        var v9 string = ("hello" + v3)
        _ = fi8
        _ = v9
    }
    switch v3 {
    case "":
        v10 := GenS1{F0: 2147483648}
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 9; fi11++ {
            v3 = v3
            _ = fi11
        }
    default:
        switch (v3 + (v3 + "\n\t")) {
        case "":
            if (190 <= 2) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        case "foo bar":
            delete(v1, 470)
        default:
            var v12 string = (v3 + v3)
            _ = v12
        }
    }
    switch 127 {
    case 9223372036854775807:
        v13 := GenS2{F0: v4, F1: (-(v4)), F2: -250}
        _ = v13
    case -525:
        fmt.Println(v3, v4)
    }
    switch (v3 + v3) {
    case "":
        fmt.Println(v3)
    case "hello":
        for fi14 := 0; fi14 < 7; fi14++ {
            v4 -= v4
            var v15 int32 = v4
            var v16 int16 = -252
            _ = fi14
            _ = v15
            _ = v16
        }
    case "foo bar":
        v4 += (v4 << ((2) & 63))
    }
    var v17 int8 = 61
    switch "a" {
    case "hello":
        v4 = v4
    }
    var v18 uint16 = 2
    var v19 bool = true
    if v19 {
        var v20 float64 = 1e10
        _ = v20
    }
    switch (v3 + (v3 + "")) {
    case "foo bar":
        v3 = (v3 + (v3 + v3))
    case "":
        var v21 int8 = (v17 ^ (v17 + 100))
        _ = v21
    case "\n\t":
        v1[1] = uint32(v18)
    default:
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 uint16 = 346
            _ = fi22
            _ = v23
        }
    }
    delete(v1, 1)
    for fi24 := 0; fi24 < 13; fi24++ {
        v2 = append(v2, v3)
        _ = fi24
    }
    var v25 uint16 = v18
    switch 39 {
    case 4294967295:
        _ = len(v1)
    case 857:
        if (2 <= 1) {
            v26 := GenS2{F0: v4, F1: -504, F2: 913}
            v19 = ((v19 && (((float64(0.0) * -2.25) < -2.25) || v19)) || (659 <= 780))
            fmt.Println(v3, v19, v17)
            _ = v26
        }
    case 9223372036854775807:
        if (0 > genFn0(v19)) {
            if (!(false || true)) {
                switch 2147483647 {
                case 684:
                    for fi27 := 0; fi27 < 11; fi27++ {
                        for fi28 := 0; fi28 < 17; fi28++ {
                            var v29 uint64 = 0
                            _ = fi28
                            _ = v29
                        }
                        fmt.Println(v4)
                        for fi30 := 0; fi30 < 10; fi30++ {
                            var v31 uint64 = 65
                            _ = fi30
                            _ = v31
                        }
                        _ = fi27
                    }
                default:
                    if (127 != 255) {
                        var v32 float64 = 3.14159
                        _ = v32
                    } else {
                        fmt.Println(v25, v25, v4)
                    }
                }
            }
        } else {
            for fi33 := 0; fi33 < 15; fi33++ {
                v3 = ""
                _ = fi33
            }
            fmt.Println(v25)
            v4 = (v4 / (((v4 + v4)) | 1))
        }
    default:
        v18 = 345
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        var v35 uint8 = 2
        fmt.Println(v17)
        for fi36 := 0; fi36 < 13; fi36++ {
            v35 -= (v35 - (^(v35)))
            _ = fi36
        }
        _ = fi34
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}

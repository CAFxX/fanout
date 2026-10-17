package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint32, p1 int16, p2 int32) int16 {
    for fi5 := 0; fi5 < 3; fi5++ {
        p0 = 4294967295
        switch fi5 {
        case -1:
            for fi6 := 0; fi6 < 4; fi6++ {
                p2 -= p2
                v7 := []int{(-(fi6)), 2147483647, fi6}
                _ = fi6
                _ = v7
            }
        case 32768:
            switch 4294967295 {
            case 127:
                p2 = p2
            case 4294967295:
                p1 += (p1 & (p1 | p1))
            default:
                p0 = (p0 << ((65535) & 63))
            }
        case -28:
            v8 := []int{}
            _ = v8
        }
        if false {
            v9 := GenS1{F0: p2}
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 int = (fi5 & -2147483648)
                var v12 string = "a"
                _ = fi10
                _ = v11
                _ = v12
            }
            fmt.Println(fi5, p0, p1)
            _ = v9
        }
        _ = fi5
    }
    v13 := GenS1{F0: p2}
    for fi14 := 0; fi14 < 10; fi14++ {
        p0 = p0
        v15 := GenS1{F0: (p2 / ((p2) | 1))}
        _ = fi14
        _ = v15
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -158
}

func main() {
    var v1 float32 = 1e10
    var v2 bool = false
    var v3 float64 = (0.1 * 0.0)
    var v4 uint64 = 18446744073709551615
    v3 -= ((v3 * v3) + v3)
    v4 = 982
    v16 := []int{}
    v2 = (!false)
    if ((((-1 > 1) && ((65535 == 65535) && ((1 != -9223372036854775808) && (4294967295 < 4294967295)))) && true) && v2) {
        v4 = (v4 * (v4 % ((0) | 1)))
    }
    if v2 {
        if false {
            var v17 int64 = 804
            switch 32768 {
            case 2147483647:
                v1 = v1
            case 9223372036854775807:
                var v18 int32 = -2147483648
                _ = v18
            }
            _ = v17
        } else {
            var v19 uint8 = 231
            var v20 int8 = -128
            v21 := GenS1{F0: -2147483648}
            _ = v19
            _ = v20
            _ = v21
        }
        v22 := GenS2{F0: 42}
        _ = v22
    }
    v23 := GenS1{F0: 2147483647}
    var v24 int16 = 1
    v25 := GenS3{F0: 1}
    _ = len(v16)
    var v26 string = (((("a" + ("a" + ((("foo bar" + "hello") + "a") + "\n\t"))) + "a") + ("\n\t" + ("hello" + ("hello" + "")))) + "foo bar")
    if (v2 && ((!v2) && (0 == 128))) {
        if false {
            var v27 int16 = 32767
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            for fi28 := 0; fi28 < 13; fi28++ {
                var v29 int64 = 0
                var v30 bool = (0 >= 255)
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = v27
        } else {
            if ((-1 != -530) || (!(v2 || false))) {
                var v31 uint32 = 0
                _ = v31
            }
            v3 += float64(v1)
        }
    } else {
        v24 += (v24 << ((375) & 63))
    }
    var v32 uint64 = (v4 >> ((75) & 63))
    v3 = float64(v1)
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v4
}

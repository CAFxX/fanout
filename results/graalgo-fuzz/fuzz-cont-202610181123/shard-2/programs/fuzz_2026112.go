package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint8) int {
    var v4 string = (((("foo bar" + "a") + "\n\t") + "") + "a")
    v5 := GenS1{F0: false}
    if (-9223372036854775808 != 2147483647) {
        switch ("\n\t" + "") {
        case "":
            switch "foo bar" {
            case "a":
                p0 = 127
            case "hello":
                p1 -= p0
            default:
                var v6 float64 = 0.0
                _ = v6
            }
        }
        v7 := GenS1{F0: (65535 >= 2)}
        var v8 int8 = -128
        _ = v7
        _ = v8
    } else {
        p1 -= (p1 << ((42) & 63))
        fmt.Println(p1, v4)
    }
    p0 = 255
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 32768
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []int{127, 0, 805, 42}
    var v3 float64 = 1.5
    var v9 string = ("foo bar" + ("foo bar" + (("a" + "\n\t") + "hello")))
    for fi10 := 0; fi10 < 20; fi10++ {
        switch v9 {
        case "hello":
            switch "\n\t" {
            case "\n\t":
                v1 = 2
            }
        case "foo bar":
            v3 = ((v3 + ((float64(float32((v3 - (v3 / 3.14159)))) + v3) + (float64(1e10) * v3))) + v3)
        }
        var v11 string = v9
        _ = fi10
        _ = v11
    }
    var v12 uint64 = 0
    var v13 uint64 = (v12 | 1)
    var v14 float32 = (float32((float64((float32(v3) + (0.1 + -2.25))) * v3)) * 0.1)
    switch v9 {
    case "":
        _ = len(v2)
    case "hello":
        switch 127 {
        case 4294967295:
            v3 -= v3
        }
    case "foo bar":
        for fi15 := 0; fi15 < 18; fi15++ {
            if (false || (2147483647 > -2147483648)) {
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 string = (v9 + ("a" + (v9 + "foo bar")))
                    _ = len(v2)
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi15
        }
    }
    if (true || ((269 == v1) && false)) {
        v1 = (v1 / ((v1) | 1))
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v2 = append(v2, -732)
    }
    var v18 int32 = -412
    var v19 int = 1
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := GenS1{F0: true}
        _ = fi20
        _ = v21
    }
    v22 := GenS2{F0: 834}
    var v23 uint32 = v1
    for fi24 := 0; fi24 < 5; fi24++ {
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 int32 = v18
            _ = fi25
            _ = v26
        }
        fmt.Println(v3, v19)
        _ = fi24
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        for fi28 := 0; fi28 < 3; fi28++ {
            var v29 uint16 = uint16(v1)
            v13 -= v12
            _ = fi28
            _ = v29
        }
        var v30 uint64 = (v13 & (v12 | v12))
        _ = fi27
        _ = v30
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v9
}

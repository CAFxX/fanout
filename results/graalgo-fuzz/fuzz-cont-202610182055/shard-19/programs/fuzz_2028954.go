package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint64, p1 int16, p2 bool) uint32 {
    v6 := GenS1{F0: 127}
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{(fi7 >> ((18446744073709551615) & 63)), fi7, 32768}
        v9 := GenS1{F0: (-(fi7))}
        switch "foo bar" {
        case "foo bar":
            _ = len(v8)
        case "hello":
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        case "":
            p0 += p0
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 719
    if p2 {
        var v11 float64 = 1e10
        var v12 int16 = -893
        v13 := GenS1{F0: 127}
        _ = v11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return v10
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint8 = 0
    var v3 uint8 = 2
    var v4 uint64 = 268
    var v5 int64 = 9223372036854775807
    v2 = (^(v3))
    var v14 float32 = -2.25
    var v15 int = -1
    switch (v15 | v15) {
    case 360:
        v4 = (v4 & v4)
    case 42:
        v16 := []int{(v15 >> ((969) & 63)), v15, (v15 ^ -139)}
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        fmt.Println(v3, v4)
        v1 = v1
        v4 -= 1
        _ = fi17
    }
    var v18 int64 = v5
    v1 += v1
    v3 = v2
    v2 = (v3 >> ((865) & 63))
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 string = ((((("" + "a") + "") + ("foo bar" + ("foo bar" + "a"))) + "hello") + "\n\t")
        switch (fi19 | fi19) {
        case 32767:
            v18 -= int64(v4)
        case 32768:
            if true {
                var v21 int64 = (v5 ^ v5)
                _ = v21
            }
        default:
            fmt.Println(v2, fi19, v1)
        }
        if (!false) {
            for fi22 := 0; fi22 < 13; fi22++ {
                switch v20 {
                case "a":
                    fmt.Println(v4)
                case "\n\t":
                    var v23 int64 = v18
                    _ = v23
                default:
                    switch 0 {
                    case 832:
                        var v24 uint64 = (v4 >> ((375) & 63))
                        _ = v24
                    }
                }
                _ = fi22
            }
        } else {
            v15 -= (fi19 & fi19)
            v25 := []int{(fi19 >> ((255) & 63)), fi19}
            fmt.Println(v4, v14)
            _ = v25
        }
        _ = fi19
        _ = v20
    }
    v14 += v14
    v2 -= 127
    var v26 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}

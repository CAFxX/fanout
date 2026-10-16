package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0(p0 string) int64 {
    fmt.Println(p0)
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return -187
}

func main() {
    var v1 bool = true
    var v2 int16 = 0
    var v3 uint8 = 245
    var v4 bool = (255 >= 42)
    v5 := []string{((("hello" + "") + "hello") + ""), ("" + ""), "a"}
    var v6 float64 = ((float64(float32((1e10 * ((-2.25 - 1e10) / 1e10)))) - 3.14159) / 1.5)
    var v7 bool = (!v1)
    var v8 float32 = (-2.25 / 1.0)
    for fi9 := 0; fi9 < 18; fi9++ {
        v4 = v1
        _ = fi9
    }
    var v10 int64 = 9223372036854775807
    fmt.Println(v3)
    fmt.Println(v7, v7)
    if len(v5) > 0 {
        v5[0] = "a"
    }
    v4 = (!(false || (42 < 330)))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 -= v6
    v11 := GenS1{F0: 798, F1: uint16(v3)}
    for fi12 := 0; fi12 < 14; fi12++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v13 uint64 = 18446744073709551615
        switch ("hello" + "\n\t") {
        case "":
            switch fi12 {
            case 552:
                var v14 float64 = 3.14159
                _ = v14
            case 9223372036854775807:
                fmt.Println(v3, v8)
            case -844:
                for fi15 := 0; fi15 < 12; fi15++ {
                    v1 = false
                    v6 += 0.1
                    v1 = false
                    _ = fi15
                }
            default:
                v5 = append(v5, "")
            }
        case "hello":
            v16 := GenS1{F0: 2, F1: 653}
            _ = v16
        case "a":
            v7 = ((!((v2 & 0) != int16(fi12))) || v4)
        }
        _ = fi12
        _ = v13
    }
    var v17 string = ("" + ("\n\t" + ("" + "foo bar")))
    for fi18 := 0; fi18 < 16; fi18++ {
        if (v6 >= 3.14159) {
            for fi19 := 0; fi19 < 3; fi19++ {
                if (42 > 2) {
                    fmt.Println(fi19)
                    if len(v5) > 0 {
                        v5[0] = v17
                    }
                } else {
                    fmt.Println(fi18, v17)
                    v7 = (!false)
                }
                _ = fi19
            }
            var v20 int16 = -705
            _ = v20
        }
        _ = fi18
    }
    var v21 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

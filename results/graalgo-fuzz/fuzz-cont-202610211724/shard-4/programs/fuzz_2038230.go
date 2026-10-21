package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 bool = false
    var v2 float32 = (float32(float64((1.5 * (1e10 * (0.0 - 0.1))))) - 0.1)
    var v3 uint32 = 0
    v4 := []string{}
    var v5 float64 = (1.5 / ((float64(-2.25) / (3.14159 - (0.0 - (1e10 - 3.14159)))) - 1.5))
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(v2, v1)
            v4 = append(v4, "a")
            _ = fi7
        }
        v1 = (v1 && (2 == 127))
        _ = fi6
    }
    var v8 uint8 = 0
    switch -928 {
    case 127:
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 string = ("\n\t" + ("" + "hello"))
            _ = fi9
            _ = v10
        }
    default:
        v4 = append(v4, (("hello" + "\n\t") + ("\n\t" + "foo bar")))
    }
    v11 := GenS1{F0: 2147483647}
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            v5 = (((float64((v2 - ((v2 * v2) - v2))) / (v5 * v5)) + v5) + ((3.14159 / v5) - (v5 - v5)))
            v2 = (((float32(-2.25) / float32(v5)) - -2.25) / v2)
            v8 = (v8 | (v8 % ((v8) | 1)))
            _ = fi13
        }
        _ = fi12
    }
    var v14 float32 = 1.5
    v15 := GenS1{F0: 272}
    var v16 int16 = -1
    v4 = append(v4, "\n\t")
    _ = len(v4)
    switch 0 {
    case -46:
        fmt.Println(v5)
    }
    if false {
        for fi17 := 0; fi17 < 13; fi17++ {
            v3 += 469
            for fi18 := 0; fi18 < 13; fi18++ {
                var v19 int16 = (v16 / (((v16 * v16)) | 1))
                v4 = append(v4, "a")
                _ = fi18
                _ = v19
            }
            v1 = (v1 && v1)
            _ = fi17
        }
        var v20 uint = 4294967295
        _ = v20
    }
    v4 = append(v4, ("a" + "a"))
    var v21 int32 = -2147483648
    var v22 uint16 = 65535
    var v23 int = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}

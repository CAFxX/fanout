package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int16 = 210
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    var v4 uint8 = 128
    var v5 int32 = -528
    v4 = v4
    v6 := GenS1{F0: ("\n\t" + "\n\t"), F1: v5}
    switch -432 {
    case 32768:
        for fi7 := 0; fi7 < 1; fi7++ {
            if (-333 >= v5) {
                v8 := GenS1{F0: "a", F1: 1}
                delete(v3, 988)
                _ = v8
            }
            _ = fi7
        }
    }
    switch 42 {
    case -233:
        if (!false) {
            for fi9 := 0; fi9 < 19; fi9++ {
                v5 = -658
                _ = fi9
            }
        } else {
            v4 = v4
        }
    case 42:
        for fi10 := 0; fi10 < 2; fi10++ {
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 string = ("" + "a")
                switch (fi11 >> ((1) & 63)) {
                case 4294967295:
                    v13 := GenS2{F0: true}
                    _ = v13
                }
                v12 = "a"
                if true {
                    if true {
                        v5 -= (v5 + (v5 >> ((382) & 63)))
                        var v14 int64 = 695
                        fmt.Println(v14, v1, v12)
                        _ = v14
                    }
                    v12 = ""
                }
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    }
    v15 := GenS1{F0: "a", F1: v5}
    fmt.Println(v4, v1)
    var v16 float64 = (3.14159 - -2.25)
    _ = len(v3)
    var v17 int64 = 9223372036854775807
    var v18 int8 = -128
    v4 = v4
    v5 = v5
    if (!(false || (v4 >= 127))) {
        v18 = 127
    } else {
        for fi19 := 0; fi19 < 18; fi19++ {
            fmt.Println(v4, v17, fi19)
            var v20 float64 = ((1.5 * (v16 / v16)) / ((0.1 - (float64((1.5 * (1.5 * 3.14159))) / v16)) / v16))
            var v21 float32 = 3.14159
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v22 := GenS1{F0: "foo bar", F1: (v5 + -2147483648)}
    fmt.Println(v18)
    v23 := v2[258]
    v3[42] = 0
    var v24 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

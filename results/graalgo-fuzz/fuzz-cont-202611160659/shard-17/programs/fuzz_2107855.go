package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ("\n\t" + "")
    var v3 int64 = -1
    var v4 uint32 = 65535
    var v5 float64 = float64((3.14159 / ((1e10 / ((1e10 / ((0.1 + (1e10 - (float32(3.14159) * float32(0.1)))) - (0.1 * 1.5))) + float32(0.0))) / -2.25)))
    v6 := []int{-307, -484, 32768, -414}
    v3 -= 9223372036854775807
    var v7 float32 = -2.25
    switch ((("" + (("foo bar" + "a") + "")) + v2) + (v2 + v2)) {
    case "foo bar":
        v8 := GenS1{F0: 0}
        _ = v8
    case "\n\t":
        if false {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            delete(v1, 127)
            v9 := GenS1{F0: 0}
            _ = v9
        }
    default:
        if false {
            switch -64 {
            case 2147483647:
                if true {
                    for fi10 := 0; fi10 < 7; fi10++ {
                        var v11 int32 = 365
                        _ = fi10
                        _ = v11
                    }
                    for fi12 := 0; fi12 < 16; fi12++ {
                        switch (("foo bar" + (v2 + (v2 + v2))) + (("foo bar" + v2) + ((v2 + (v2 + ("foo bar" + ((v2 + (v2 + v2)) + (v2 + v2))))) + (v2 + v2)))) {
                        case "\n\t":
                            v5 -= ((v5 - (v5 - v5)) + 1.5)
                        case "":
                            var v13 bool = false
                            _ = v13
                        case "foo bar":
                            var v14 float64 = ((0.1 - 1e10) - v5)
                            _ = v14
                        }
                        if true {
                            var v15 int = fi12
                            _ = v15
                        } else {
                            fmt.Println(v4)
                            var v16 int8 = -13
                            v7 -= (v7 + (v7 + v7))
                            _ = v16
                        }
                        _ = fi12
                    }
                }
            }
            if true {
                v17 := GenS2{F0: 42, F1: 104}
                _ = v17
            } else {
                v3 = v3
            }
        }
    }
    var v18 int16 = 281
    for fi19 := 0; fi19 < 15; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            v6 = append(v6, fi19)
            _ = fi20
        }
        _ = fi19
    }
    var v21 int16 = 32767
    v22 := v1[-2147483648]
    var v23 int16 = -851
    v4 -= (v22 & (v22 - v4))
    var v24 uint8 = 255
    for fi25 := 0; fi25 < 15; fi25++ {
        var v26 bool = ((v3 / ((-1) | 1)) < (v3 / ((v3) | 1)))
        var v27 uint32 = (v4 | 747)
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

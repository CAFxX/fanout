package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 float64 = ((((0.0 - ((-2.25 * 0.0) - 0.1)) * (0.0 / (0.1 / 1.0))) + float64(0.1)) + 0.0)
    v2 := make(map[string]uint32)
    var v3 float64 = -2.25
    var v4 int32 = 1
    v5 := make(map[string]string)
    fmt.Println(v3, v1)
    var v6 int16 = -684
    v7 := GenS1{F0: true}
    v8 := GenS1{F0: (65535 <= 42)}
    v1 -= ((v3 + v1) * (v3 / 1.5))
    delete(v2, ("a" + ""))
    v9 := GenS2{F0: -24}
    var v10 uint = 18446744073709551615
    switch ("\n\t" + (((("hello" + "\n\t") + "\n\t") + "\n\t") + ("a" + ""))) {
    case "a":
        var v11 string = "a"
        _ = v11
    }
    if ((v1 - (v3 - -2.25)) >= 0.0) {
        fmt.Println(v1, v10)
    } else {
        switch 9223372036854775807 {
        case 2147483647:
            switch 0 {
            case -641:
                if ((v4 / ((v4) | 1)) <= 42) {
                    switch (((("\n\t" + ((("" + ("" + ("a" + ""))) + "foo bar") + "a")) + "\n\t") + (("" + "foo bar") + "\n\t")) + ("a" + "\n\t")) {
                    case "\n\t":
                        v3 = v1
                    case "a":
                        var v12 float64 = -2.25
                        _ = v12
                    case "hello":
                        v6 = v6
                    }
                    if (((true || (!((1e10 / (v1 / ((v1 + (((1e10 * (float64(3.14159) * v3)) * (1.5 * v1)) * v3)) * ((float64(float32((v1 / ((((v1 / (v3 - 1.5)) + ((float64(3.14159) + v3) + v3)) + (v3 + (v3 - float64(0.1)))) + v3)))) + (1.5 + ((1.5 + v1) - float64(1.5)))) * (0.0 - 1e10))))) == (v1 / v3)))) || false) || (v10 < v10)) {
                        for fi13 := 0; fi13 < 7; fi13++ {
                            v14 := []int{(fi13 / ((fi13) | 1)), int(v4)}
                            _ = fi13
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 20; fi15++ {
                            fmt.Println(v4, v4)
                            if true {
                                v3 -= v1
                            } else {
                                var v16 int32 = v4
                                var v17 bool = (fi15 != (fi15 << ((v10) & 63)))
                                v18 := GenS3{F0: v17, F1: (v4 | v4)}
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            _ = fi15
                        }
                        var v19 int16 = (v6 << ((721) & 63))
                        _ = v19
                    }
                }
            }
        case 32767:
            v4 += (v4 << (((v10 - v10)) & 63))
        case 128:
            v20 := GenS4{F0: v4}
            _ = v20
        }
        fmt.Println(v3, v3, v1)
        v6 += -726
    }
    v6 = -383
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

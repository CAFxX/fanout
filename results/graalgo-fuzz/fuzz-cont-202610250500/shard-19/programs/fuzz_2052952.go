package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[int]string)
    v2 := []int{}
    var v3 int64 = 42
    var v4 string = "\n\t"
    var v5 int16 = 1000
    var v6 string = "hello"
    switch (v6 + "hello") {
    case "\n\t":
        var v7 bool = (-2.25 <= 1e10)
        _ = v7
    default:
        v4 = v4
    }
    v8 := GenS1{F0: v3}
    var v9 bool = ((2147483648 == 1) || true)
    v9 = false
    v10 := GenS2{F0: 0, F1: (1e10 != 3.14159)}
    switch -1 {
    case 323:
        v11 := GenS1{F0: (v3 << ((255) & 63))}
        _ = v11
    }
    v1[-511] = "hello"
    v4 = "hello"
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 uint8 = 255
        if false {
            v4 = v6
        } else {
            if (!v9) {
                if v9 {
                    v4 = ("hello" + v6)
                    fmt.Println(v6, fi12)
                    for fi14 := 0; fi14 < 20; fi14++ {
                        var v15 int64 = int64(v13)
                        v13 += v13
                        _ = fi14
                        _ = v15
                    }
                }
                if false {
                    var v16 uint32 = 2147483648
                    switch (("" + v6) + (v4 + "\n\t")) {
                    case "hello":
                        v17 := v1[fi12]
                        _ = v17
                    default:
                        v18 := GenS2{F0: -1, F1: (v13 <= v13)}
                        _ = v18
                    }
                    _ = v16
                } else {
                    switch (v4 + v6) {
                    case "a":
                        var v19 uint32 = 836
                        _ = v19
                    default:
                        v4 = "foo bar"
                    }
                }
            }
            v6 = ("foo bar" + v6)
            _ = len(v2)
        }
        var v20 uint16 = 480
        _ = fi12
        _ = v13
        _ = v20
    }
    _ = len(v1)
    switch 32768 {
    case 4294967295:
        for fi21 := 0; fi21 < 12; fi21++ {
            for fi22 := 0; fi22 < 17; fi22++ {
                v1[fi21] = (v6 + ("\n\t" + "\n\t"))
                _ = fi22
            }
            switch fi21 {
            case 0:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                for fi23 := 0; fi23 < 5; fi23++ {
                    fmt.Println(fi23, v5, fi21)
                    v24 := GenS2{F0: 1, F1: (1.5 == ((1e10 / (1e10 + 0.0)) + 0.0))}
                    fmt.Println(v3)
                    _ = fi23
                    _ = v24
                }
            }
            v25 := GenS2{F0: 42, F1: true}
            _ = fi21
            _ = v25
        }
    default:
        var v26 int = int(v5)
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}

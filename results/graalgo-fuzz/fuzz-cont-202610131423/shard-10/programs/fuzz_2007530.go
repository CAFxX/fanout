package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 uint32; F2 int16 }

func main() {
    v1 := []string{}
    var v2 bool = true
    var v3 uint32 = 2147483648
    var v4 uint32 = (v3 | v3)
    var v5 bool = false
    var v6 bool = (("\n\t" + ("hello" + "")) != "hello")
    var v7 uint64 = 2
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    v8 := GenS1{F0: v3, F1: 42}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v9 float32 = (0.1 + -2.25)
    var v10 float32 = 3.14159
    var v11 string = ("a" + ("a" + "\n\t"))
    var v12 uint8 = 128
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            if true {
                switch 9223372036854775807 {
                case 127:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case -2147483648:
                    v9 = float32(0.0)
                    if true {
                        if len(v1) > 0 {
                            v1[0] = v11
                        }
                    }
                case 501:
                    if v5 {
                        fmt.Println(fi13, v2, fi14)
                        v4 += uint32(v12)
                    }
                }
                v7 -= (v7 ^ (v7 * v7))
            }
            v6 = (42 >= 0)
            fmt.Println(v7)
            if v5 {
                v2 = v6
                if false {
                    v11 = (("hello" + "a") + v11)
                    v11 = (v11 + (v11 + "hello"))
                    fmt.Println(v11, v4)
                } else {
                    v15 := GenS2{F0: -2147483648}
                    _ = v15
                }
                _ = len(v1)
            }
            _ = fi14
        }
        v2 = (uint32(v12) == v3)
        _ = fi13
    }
    if v2 {
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 uint8 = (v12 * (v12 - 128))
            v1 = append(v1, "a")
            for fi18 := 0; fi18 < 12; fi18++ {
                _ = len(v1)
                if true {
                    v5 = (255 == 2)
                    v19 := GenS3{F0: 303, F1: 42, F2: 171}
                    _ = v19
                }
                v7 -= v7
                var v20 uint32 = (v3 | v4)
                _ = fi18
                _ = v20
            }
            _ = fi16
            _ = v17
        }
    }
    var v21 int64 = 1
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 bool = (!v5)
        v2 = (true && (!((0 < -128) && true)))
        _ = fi22
        _ = v23
    }
    switch 32767 {
    case -9223372036854775808:
        switch ("\n\t" + v11) {
        case "a":
            var v24 int = 32767
            _ = v24
        case "hello":
            if (false || ((!v6) || v6)) {
                var v25 uint64 = v7
                var v26 string = v11
                _ = v25
                _ = v26
            }
        case "foo bar":
            v11 = v11
        }
    case 2147483647:
        v2 = true
    case 4294967295:
        v9 = (1.5 * float32((((float64(float32(0.0)) + -2.25) * (-2.25 / 0.1)) - -2.25)))
    }
    var v27 bool = (false && true)
    var v28 int32 = 42
    v11 = v11
    v29 := GenS3{F0: (-(v21)), F1: (v4 & 2147483648), F2: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

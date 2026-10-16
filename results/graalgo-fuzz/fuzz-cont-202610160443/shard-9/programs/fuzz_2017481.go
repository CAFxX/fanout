package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint32 }
func genFn0() int {
    var v7 uint8 = 196
    v8 := GenS1{F0: 355, F1: 2147483647, F2: (("\n\t" + "a") + "hello")}
    switch 32767 {
    case 9223372036854775807:
        var v9 int32 = 2147483647
        _ = v9
    case 2147483647:
        var v10 string = "a"
        _ = v10
    }
    var v11 uint64 = 18446744073709551615
    _ = v11
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 3.14159
    v3 := make(map[string]bool)
    v4 := make(map[string]string)
    var v5 int8 = 127
    v6 := make(map[uint32]uint32)
    switch (("" + ("foo bar" + ("\n\t" + "\n\t"))) + "foo bar") {
    case "hello":
        fmt.Println(v2)
    case "":
        if (4294967295 >= 2147483647) {
            v12 := []int{128}
            _ = v12
        } else {
            switch "" {
            case "hello":
                var v13 int = 32768
                _ = v13
            default:
                switch (((((("a" + "foo bar") + ("\n\t" + "")) + (("a" + "foo bar") + ("\n\t" + "\n\t"))) + "foo bar") + "hello") + "a") {
                case "hello":
                    v14 := GenS2{F0: -567}
                    _ = v14
                default:
                    v15 := GenS3{F0: 2, F1: ""}
                    _ = v15
                }
            }
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 uint16 = 255
                v5 -= 127
                for fi18 := 0; fi18 < 17; fi18++ {
                    v19 := []int{(fi16 >> ((2) & 63)), fi16, (fi16 - fi16)}
                    _ = fi18
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
            fmt.Println(v2, v2)
        }
    case "foo bar":
        fmt.Println(v5, v5)
    }
    v5 = 127
    var v20 float32 = 1.5
    var v21 float64 = 3.14159
    switch -824 {
    case 754:
        for fi22 := 0; fi22 < 10; fi22++ {
            for fi23 := 0; fi23 < 1; fi23++ {
                if false {
                    if true {
                        var v24 uint = 18446744073709551615
                        delete(v3, "\n\t")
                        _ = v24
                    }
                    var v25 int8 = (-(v5))
                    v26 := []int{(fi23 - (-(fi22))), fi22, fi23, (fi23 >> ((2) & 63))}
                    _ = v25
                    _ = v26
                }
                v27 := []int{}
                if (0 <= 887) {
                    var v28 float32 = v20
                    _ = v28
                } else {
                    var v29 float64 = 0.1
                    var v30 string = ("" + "\n\t")
                    _ = v29
                    _ = v30
                }
                fmt.Println(v20, v20)
                _ = fi23
                _ = v27
            }
            v31 := GenS4{F0: 432}
            _ = fi22
            _ = v31
        }
    case -574:
        v21 = 0.0
    case 100:
        fmt.Println(v2, v5)
    }
    if true {
        for fi32 := 0; fi32 < 10; fi32++ {
            var v33 uint16 = 32767
            _ = fi32
            _ = v33
        }
    } else {
        var v34 uint = 1
        v20 += (((v20 * 0.1) * -2.25) - (0.1 - ((0.0 * float32(float64(1.5))) / (1e10 - float32(v21)))))
        _ = v34
    }
    v35 := GenS2{F0: -131}
    var v36 uint = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}

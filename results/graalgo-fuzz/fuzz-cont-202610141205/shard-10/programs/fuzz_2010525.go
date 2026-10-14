package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int64; F1 int }

func main() {
    v1 := []int{4294967295, -1000}
    var v2 float64 = 1.5
    v3 := []string{"hello", ("foo bar" + ("foo bar" + ("" + "\n\t"))), "a"}
    if len(v3) > 0 {
        v3[0] = (("hello" + ("foo bar" + "foo bar")) + "\n\t")
    }
    if true {
        v1 = append(v1, 32768)
        v2 += v2
        for fi4 := 0; fi4 < 14; fi4++ {
            for fi5 := 0; fi5 < 13; fi5++ {
                var v6 uint = 255
                v6 = (v6 + v6)
                var v7 bool = false
                switch (fi5 - fi5) {
                case 32767:
                    fmt.Println(fi5, v2)
                }
                switch ("\n\t" + "hello") {
                case "hello":
                    fmt.Println(v6)
                case "\n\t":
                    v8 := GenS1{F0: (^(v6)), F1: "\n\t"}
                    _ = v8
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    v1 = append(v1, (fi5 ^ 800))
                }
                _ = fi5
                _ = v6
                _ = v7
            }
            _ = fi4
        }
    } else {
        var v9 float64 = v2
        _ = v9
    }
    switch "a" {
    case "\n\t":
        v10 := GenS1{F0: 4294967295, F1: (("\n\t" + "hello") + ("" + "\n\t"))}
        _ = v10
    case "foo bar":
        v11 := GenS2{F0: 9223372036854775807, F1: 2147483647}
        _ = v11
    case "":
        v2 = (v2 * float64(float32(float64((0.1 / 1.0)))))
    default:
        if (3.14159 != float32((1.5 * v2))) {
            var v12 float64 = ((float64(0.0) - (0.1 - 1.5)) / 1.5)
            _ = v12
        }
    }
    var v13 string = "foo bar"
    var v14 bool = true
    v1 = append(v1, 32767)
    var v15 int32 = -1
    if (76 <= 82) {
        var v16 uint32 = 4294967295
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int64; F2 int32 }
func genFn0(p0 bool, p1 uint16) int8 {
    var v6 int16 = 1000
    v7 := GenS1{F0: 0, F1: -236}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := []string{"hello", "foo bar", ("hello" + ("a" + "hello")), "a"}
    var v2 int = 698
    var v3 int64 = 0
    v4 := []int{348, (-(v2)), v2}
    var v5 float64 = -2.25
    var v8 int = v2
    v3 += v3
    _ = len(v4)
    for fi9 := 0; fi9 < 13; fi9++ {
        if false {
            var v10 float64 = (v5 * -2.25)
            v1 = append(v1, "")
            v5 = (v10 / ((v10 * float64(3.14159)) + (float64(1e10) - -2.25)))
            _ = v10
        } else {
            v3 = v3
            v11 := GenS1{F0: 2, F1: fi9}
            v12 := GenS2{F0: genFn0(false, 0)}
            _ = v11
            _ = v12
        }
        if false {
            switch "a" {
            case "\n\t":
                var v13 uint = 2
                _ = v13
            case "foo bar":
                var v14 uint = 525
                _ = v14
            case "":
                v15 := GenS3{F0: 1000, F1: -698, F2: 0}
                _ = v15
            }
            v8 = 128
            var v16 int32 = 238
            _ = v16
        } else {
            switch (v8 ^ 128) {
            case 42:
                for fi17 := 0; fi17 < 15; fi17++ {
                    if ((936 < 2147483647) || false) {
                        var v18 float32 = 0.0
                        fmt.Println(v18, v8, v5)
                        _ = v18
                    }
                    _ = fi17
                }
            case 2147483647:
                v5 += 1e10
            case 1:
                if false {
                    v2 = v8
                }
            default:
                if len(v4) > 0 {
                    v4[0] = 9223372036854775807
                }
            }
            v3 = v3
        }
        v5 = (0.0 / v5)
        _ = fi9
    }
    var v19 int16 = 760
    v2 += (v2 ^ v2)
    v2 = (v2 & 127)
    var v20 uint = 0
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}

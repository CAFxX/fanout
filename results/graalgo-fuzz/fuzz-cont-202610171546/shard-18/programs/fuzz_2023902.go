package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 bool }

func main() {
    v1 := []string{("hello" + "\n\t"), "a", "hello"}
    var v2 float64 = float64(-2.25)
    var v3 uint16 = 346
    v4 := make(map[int]string)
    var v5 int64 = 0
    v3 += (v3 + (v3 >> ((944) & 63)))
    v6 := GenS1{F0: 2, F1: -49, F2: 0}
    v7 := GenS1{F0: (v3 / (((v3 & (v3 * 396))) | 1)), F1: 48, F2: -1}
    v5 += v5
    if (2147483648 >= 761) {
        fmt.Println(v3)
    }
    v4[-1] = "foo bar"
    if false {
        v3 -= 0
        if (255 == 42) {
            if (("foo bar" + "a") != ("\n\t" + "a")) {
                fmt.Println(v5)
                switch ("hello" + "foo bar") {
                case "\n\t":
                    delete(v4, 128)
                case "":
                    v8 := GenS2{F0: -825, F1: ("hello" > "")}
                    _ = v8
                default:
                    var v9 uint = 0
                    _ = v9
                }
            } else {
                var v10 int16 = 539
                if true {
                    var v11 string = "hello"
                    v3 += v3
                    _ = v11
                }
                _ = v10
            }
        }
    }
    v12 := GenS3{F0: true}
    v5 += -9223372036854775808
    v13 := GenS2{F0: -1, F1: (true || false)}
    v14 := GenS2{F0: 42, F1: (true || (65535 >= 65535))}
    if len(v1) > 0 {
        v1[0] = ("hello" + (("a" + ("a" + "\n\t")) + "foo bar"))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

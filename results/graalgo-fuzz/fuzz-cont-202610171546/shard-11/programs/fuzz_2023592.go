package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 int = 128
    v3 := []string{"\n\t"}
    var v4 uint8 = 178
    var v5 uint = 18446744073709551615
    var v6 uint = 65535
    for fi7 := 0; fi7 < 16; fi7++ {
        switch "hello" {
        case "hello":
            v2 += (fi7 * 42)
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 float32 = 0.1
            var v10 uint16 = 578
            v1["\n\t"] = "\n\t"
            v11 := GenS1{F0: (^(v5)), F1: -388}
            v3 = append(v3, ("" + (("hello" + (("foo bar" + ("" + ("foo bar" + ("a" + "\n\t")))) + "a")) + (("" + ("hello" + ("\n\t" + ""))) + ("hello" + "foo bar")))))
            switch "hello" {
            case "a":
                fmt.Println(v6, fi8)
            case "foo bar":
                if false {
                    v9 -= 1.5
                    v10 -= 32767
                    fmt.Println(v9, v4)
                    v3 = append(v3, "foo bar")
                    v2 += 0
                    var v12 bool = (!(3.14159 >= (1.5 / 0.1)))
                    var v13 float64 = float64(0.0)
                    _ = v12
                    _ = v13
                }
            case "hello":
                fmt.Println(fi7, v2, fi8)
            }
            var v14 uint32 = 42
            _ = fi8
            _ = v10
            _ = v11
            _ = v14
            _ = v9
        }
        _ = fi7
    }
    v4 -= (v4 * 42)
    v6 = uint(v4)
    v15 := GenS1{F0: v5, F1: 1}
    var v16 string = ((("a" + ("a" + "\n\t")) + "hello") + "\n\t")
    v17 := GenS1{F0: (v6 & (^(v6))), F1: -1}
    var v18 uint = (v6 >> (((v6 << (((v6 * v6)) & 63))) & 63))
    v6 -= v6
    switch v16 {
    case "":
        fmt.Println(v18, v6)
    case "\n\t":
        v18 += v18
    case "foo bar":
        delete(v1, (("a" + (v16 + ((v16 + (v16 + v16)) + "a"))) + (v16 + "foo bar")))
    }
    v18 = (v5 % ((v6) | 1))
    var v19 float32 = float32(1.5)
    v2 += v2
    if false {
        fmt.Println(v16)
    }
    v2 -= (v2 * v2)
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := GenS1{F0: (v5 % (((^(v18))) | 1)), F1: 252}
        switch v16 {
        case "":
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 int32 = 42
                v6 -= 255
                if (v16 > (v16 + (v16 + v16))) {
                    fmt.Println(v6)
                }
                _ = fi22
                _ = v23
            }
        case "\n\t":
            v16 = v16
        }
        _ = fi20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

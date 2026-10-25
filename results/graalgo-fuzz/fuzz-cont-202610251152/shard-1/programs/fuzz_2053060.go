package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 bool; F2 string }
type GenS3 struct { F0 uint32; F1 uint; F2 int16 }

func main() {
    v1 := []string{"\n\t", "foo bar", "foo bar", (("\n\t" + ((("\n\t" + "\n\t") + ("" + "a")) + "")) + ("" + "hello"))}
    v2 := make(map[string]int)
    var v3 int32 = -1
    var v4 uint32 = 0
    fmt.Println(v3, v4)
    for fi5 := 0; fi5 < 15; fi5++ {
        fmt.Println(v4, fi5, fi5)
        v6 := GenS1{F0: ("hello" <= "")}
        if false {
            var v7 int8 = -50
            fmt.Println(v4, v3)
            _ = v7
        } else {
            for fi8 := 0; fi8 < 20; fi8++ {
                v3 -= (v3 / (((v3 % ((v3) | 1))) | 1))
                _ = fi8
            }
            v4 -= (v4 >> ((4294967295) & 63))
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(v3)
    v3 = v3
    var v9 uint64 = 2
    v9 += v9
    v10 := GenS2{F0: "", F1: false, F2: "a"}
    if true {
        v11 := GenS3{F0: (v4 & v4), F1: 484, F2: 1000}
        _ = v11
    } else {
        fmt.Println(v4, v9, v9)
        for fi12 := 0; fi12 < 11; fi12++ {
            switch "" {
            case "hello":
                v4 += v4
            case "":
                for fi13 := 0; fi13 < 14; fi13++ {
                    v14 := GenS1{F0: (0.1 != ((float64(0.0) - (((((1.5 - ((3.14159 + 1e10) * 0.1)) * (3.14159 - 3.14159)) - (((0.0 - float64(1e10)) + 1e10) - float64(0.0))) * 0.0) * 0.1)) + 0.0))}
                    v4 = v4
                    _ = fi13
                    _ = v14
                }
            }
            fmt.Println(fi12)
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := GenS2{F0: ("foo bar" + "hello"), F1: true, F2: ("a" + "hello")}
                v4 = 42
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi15
                _ = v16
            }
            _ = fi12
        }
        v3 -= (v3 ^ v3)
    }
    v9 += v9
    delete(v2, "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}

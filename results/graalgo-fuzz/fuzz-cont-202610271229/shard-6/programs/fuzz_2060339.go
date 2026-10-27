package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 int; F1 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 0
    var v3 int16 = 794
    var v4 uint64 = 1
    v5 := []string{("a" + ""), ""}
    v6 := make(map[int]int)
    v2 -= v2
    fmt.Println(v3, v4)
    if (!(true || true)) {
        if false {
            switch -2147483648 {
            case 32768:
                v3 = v3
            }
            var v7 string = (("\n\t" + "\n\t") + ("hello" + "a"))
            if (4294967295 < 42) {
                v4 -= v4
                v3 += (v3 ^ 1000)
            }
            _ = v7
        }
        for fi8 := 0; fi8 < 11; fi8++ {
            _ = len(v1)
            v3 = -733
            _ = fi8
        }
        v3 -= 42
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch 128 {
    case 32767:
        for fi9 := 0; fi9 < 6; fi9++ {
            fmt.Println(fi9, fi9)
            _ = fi9
        }
    case 9223372036854775807:
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 string = "a"
            for fi12 := 0; fi12 < 7; fi12++ {
                switch ((((v11 + v11) + v11) + (v11 + "a")) + "foo bar") {
                case "a":
                    v2 = (v2 ^ v2)
                case "\n\t":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v13 bool = false
                    _ = v13
                default:
                    delete(v6, fi12)
                }
                v3 = v3
                var v14 uint32 = 2
                _ = fi12
                _ = v14
            }
            v15 := GenS1{F0: (822 == uint32(v2)), F1: 518, F2: -74}
            _ = fi10
            _ = v11
            _ = v15
        }
    case 127:
        if true {
            _ = len(v1)
        }
    }
    v5 = append(v5, "foo bar")
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 int64 = 42
        fmt.Println(v3, fi16, fi16)
        if (float64(0.1) <= 1e10) {
            v18 := GenS2{F0: fi16, F1: (v3 - v3)}
            for fi19 := 0; fi19 < 11; fi19++ {
                v6[(-(fi16))] = int(v3)
                _ = fi19
            }
            _ = v18
        } else {
            v17 = (v17 | v17)
        }
        _ = fi16
        _ = v17
    }
    var v20 string = ("a" + "\n\t")
    v21 := GenS2{F0: 779, F1: v3}
    v3 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

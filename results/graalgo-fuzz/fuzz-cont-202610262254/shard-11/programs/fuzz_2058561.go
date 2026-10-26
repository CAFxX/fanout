package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int32 }
func genFn0(p0 int32) int64 {
    p0 += 1
    v6 := GenS1{F0: 0}
    _ = p0
    _ = v6
    return 271
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = -1
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    var v5 uint = 402
    fmt.Println(v2, v2)
    switch ("" + "hello") {
    case "a":
        delete(v4, 749)
    case "":
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{(fi7 & fi7), fi7, (fi7 + fi7)}
            _ = fi7
            _ = v8
        }
    default:
        if true {
            fmt.Println(v2, v5)
            var v9 uint8 = 2
            if (true || false) {
                v9 -= 127
            }
            _ = v9
        }
    }
    var v10 uint16 = 1
    switch 32768 {
    case 128:
        if (4294967295 < 187) {
            switch "\n\t" {
            case "hello":
                var v11 bool = (!(v5 > 42))
                _ = v11
            case "":
                v2 -= (v2 ^ v2)
            case "foo bar":
                var v12 string = "foo bar"
                _ = v12
            }
            var v13 float32 = 3.14159
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 int32 = 42
            v15 = (v15 + v15)
            _ = fi14
            _ = v15
        }
    }
    switch 32767 {
    case -666:
        for fi16 := 0; fi16 < 20; fi16++ {
            v2 -= (v2 & (v2 / (((v2 ^ -1)) | 1)))
            if false {
                switch fi16 {
                case 42:
                    switch fi16 {
                    case 4294967295:
                        v17 := []int{}
                        _ = v17
                    case 42:
                        var v18 int = fi16
                        _ = v18
                    }
                }
                v2 = int16(fi16)
            }
            _ = fi16
        }
    }
    v5 = v5
    v19 := GenS2{F0: 449, F1: 0}
    var v20 uint64 = 42
    _ = len(v4)
    fmt.Println(v2, v10, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}

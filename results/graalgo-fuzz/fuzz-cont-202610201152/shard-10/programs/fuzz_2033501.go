package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint16; F2 int64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    var v3 uint64 = 2
    fmt.Println(v2, v2)
    var v4 uint32 = 44
    v3 -= 2
    if (127 < 127) {
        for fi5 := 0; fi5 < 5; fi5++ {
            switch fi5 {
            case 32767:
                v6 := []int{}
                _ = v6
            case 42:
                for fi7 := 0; fi7 < 11; fi7++ {
                    var v8 int16 = -32768
                    v9 := []int{(fi7 ^ fi7), fi5, 45, fi7}
                    var v10 int32 = 2147483647
                    _ = fi7
                    _ = v10
                    _ = v8
                    _ = v9
                }
            case 2147483647:
                var v11 int32 = 0
                _ = v11
            default:
                v2 = (!v2)
            }
            if (582 <= 50) {
                v12 := []int{}
                v2 = (!(((("\n\t" + "") + "") > "hello") && (v2 || (!(v2 && v2)))))
                _ = v12
            }
            for fi13 := 0; fi13 < 14; fi13++ {
                switch (^(fi5)) {
                case 42:
                    var v14 float32 = 0.0
                    _ = v14
                case 32767:
                    v15 := []int{}
                    _ = v15
                case 127:
                    _ = len(v1)
                default:
                    var v16 float64 = 1.5
                    _ = v16
                }
                _ = fi13
            }
            _ = fi5
        }
    }
    v1[608] = 1
    switch 755 {
    case -540:
        var v17 uint16 = 32767
        _ = v17
    case 0:
        switch ((("a" + "foo bar") + "foo bar") + "\n\t") {
        case "hello":
            var v18 int16 = -1
            _ = v18
        }
    }
    v4 -= (^(v4))
    fmt.Println(v2)
    for fi19 := 0; fi19 < 7; fi19++ {
        switch ("hello" + "\n\t") {
        case "hello":
            var v20 int16 = 32767
            _ = v20
        case "":
            fmt.Println(v4)
        default:
            v21 := GenS1{F0: -1}
            _ = v21
        }
        _ = fi19
    }
    v22 := GenS2{F0: ("foo bar" + ""), F1: 0, F2: 0}
    var v23 string = ""
    v24 := []int{111, 2147483647, 128, 0}
    v1[v4] = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}

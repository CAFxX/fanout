package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int }
type GenS2 struct { F0 bool }

func main() {
    var v1 int64 = 42
    var v2 uint = 65535
    var v3 uint8 = 128
    v4 := make(map[string]int)
    var v5 int32 = -1
    var v6 uint16 = 2
    for fi7 := 0; fi7 < 2; fi7++ {
        switch "hello" {
        case "foo bar":
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 int32 = (v5 | v5)
                _ = fi8
                _ = v9
            }
        }
        var v10 uint = v2
        var v11 uint16 = (v6 * v6)
        _ = fi7
        _ = v10
        _ = v11
    }
    if true {
        for fi12 := 0; fi12 < 13; fi12++ {
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 uint16 = (v6 ^ v6)
                if false {
                    var v15 int = fi12
                    if false {
                        v4[""] = 636
                        var v16 uint32 = 2147483647
                        _ = v16
                    }
                    var v17 int8 = 0
                    if (v15 == fi12) {
                        switch "\n\t" {
                        case "hello":
                            delete(v4, ("a" + "hello"))
                        default:
                            var v18 float64 = -2.25
                            _ = v18
                        }
                        v15 -= -269
                    } else {
                        v19 := []int{2147483647, (^(v15)), (v15 & fi12), (fi12 + fi12)}
                        var v20 uint32 = 891
                        var v21 float32 = 1.5
                        v22 := GenS1{F0: 0, F1: 1000, F2: fi13}
                        _ = v19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    _ = v15
                    _ = v17
                }
                v6 = v14
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    }
    var v23 uint64 = 42
    _ = len(v4)
    v2 -= 65535
    v2 += v2
    v24 := []int{570, 42}
    v5 += (v5 + v5)
    v2 = 65535
    switch "" {
    case "foo bar":
        var v25 uint64 = (v23 >> ((v2) & 63))
        _ = v25
    case "\n\t":
        v1 -= (v1 / ((v1) | 1))
    case "hello":
        for fi26 := 0; fi26 < 1; fi26++ {
            if (true && (0 <= 962)) {
                var v27 int64 = v1
                _ = v27
            } else {
                v28 := GenS1{F0: 0, F1: 1, F2: (fi26 / ((int(v2)) | 1))}
                v29 := GenS1{F0: 127, F1: 1, F2: 42}
                var v30 int16 = 42
                _ = v28
                _ = v29
                _ = v30
            }
            v31 := GenS2{F0: false}
            _ = fi26
            _ = v31
        }
    }
    delete(v4, "")
    var v32 uint32 = 2147483648
    if false {
        for fi33 := 0; fi33 < 6; fi33++ {
            var v34 int = fi33
            _ = fi33
            _ = v34
        }
    }
    var v35 int8 = 1
    if ((!true) || false) {
        v32 = (v32 << (((v2 << (((^(v2))) & 63))) & 63))
    }
    var v36 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}

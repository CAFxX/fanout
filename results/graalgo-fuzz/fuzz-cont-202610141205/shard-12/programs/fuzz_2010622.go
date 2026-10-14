package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
func genFn0(p0 uint16) uint8 {
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: p0}
        v9 := GenS1{F0: 0}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 42}
    v11 := []int{}
    _ = p0
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 uint32 = 251
    var v2 int64 = 0
    v3 := make(map[uint32]int)
    v4 := []string{}
    var v5 int = 538
    var v6 string = "foo bar"
    fmt.Println(v6, v2, v1)
    var v12 int16 = 42
    if (true && true) {
        fmt.Println(v2, v5, v5)
    } else {
        v4 = append(v4, "foo bar")
        v5 = 472
    }
    v6 = v6
    fmt.Println(v6)
    v3[(v1 >> ((588) & 63))] = 1
    v13 := GenS2{F0: v6}
    if false {
        var v14 int16 = int16(v5)
        _ = v14
    }
    if (v12 >= v12) {
        if false {
            switch -814 {
            case 508:
                for fi15 := 0; fi15 < 11; fi15++ {
                    if false {
                        var v16 string = v6
                        _ = v16
                    }
                    v1 = v1
                    var v17 int32 = 0
                    _ = fi15
                    _ = v17
                }
            }
            var v18 string = "\n\t"
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 string = "foo bar"
                v6 = ("\n\t" + "")
                var v21 string = (v20 + "")
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = v18
        }
    }
    v22 := GenS2{F0: v6}
    v23 := GenS1{F0: 32767}
    fmt.Println(v1)
    var v24 int64 = v2
    if true {
        fmt.Println(v12, v1, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

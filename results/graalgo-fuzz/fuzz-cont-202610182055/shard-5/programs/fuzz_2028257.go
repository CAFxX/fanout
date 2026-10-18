package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string, p1 uint32) int64 {
    p1 += p1
    v7 := []int{-1, -9223372036854775808, 2147483647, -9223372036854775808}
    var v8 uint64 = 0
    v8 = v8
    switch "hello" {
    case "foo bar":
        p0 = ((p0 + p0) + p0)
    case "a":
        if false {
            v8 = (v8 | v8)
            var v9 uint32 = p1
            var v10 bool = true
            _ = v10
            _ = v9
        } else {
            var v11 int = 189
            p0 = p0
            _ = v11
        }
    default:
        p1 = (p1 << ((0) & 63))
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 361
}

func main() {
    v1 := []int{-774, -555}
    var v2 bool = ((("\n\t" + "") + "") >= "a")
    v3 := []string{}
    v4 := make(map[int]int)
    v5 := make(map[uint32]int)
    var v6 bool = (0.0 < (0.0 + ((0.0 + 3.14159) - (3.14159 * float64((float32(1.5) / float32(((1e10 / -2.25) - float64((((3.14159 / (1.5 + 0.1)) * ((1e10 - float32(0.1)) * 1e10)) * 1.5))))))))))
    if false {
        if v2 {
            var v12 uint16 = 127
            _ = len(v5)
            _ = v12
        } else {
            v6 = v6
            for fi13 := 0; fi13 < 10; fi13++ {
                v14 := GenS1{F0: 127}
                if (!(42 == 42)) {
                    if true {
                        if (!true) {
                            switch (fi13 - fi13) {
                            case -396:
                                v15 := GenS1{F0: 0}
                                _ = v15
                            default:
                                v16 := v4[fi13]
                                _ = v16
                            }
                            var v17 float32 = 1e10
                            for fi18 := 0; fi18 < 11; fi18++ {
                                switch (fi18 & (fi13 / (((fi18 - (fi18 - fi18))) | 1))) {
                                case 944:
                                    v2 = v6
                                }
                                fmt.Println(v6)
                                if v6 {
                                    var v19 uint64 = 0
                                    _ = v19
                                }
                                _ = fi18
                            }
                            _ = v17
                        }
                        var v20 int = fi13
                        _ = v20
                    } else {
                        v6 = (-59 == 42)
                        for fi21 := 0; fi21 < 2; fi21++ {
                            v22 := GenS1{F0: 0}
                            _ = fi21
                            _ = v22
                        }
                    }
                }
                _ = fi13
                _ = v14
            }
            v6 = true
        }
        var v23 float32 = (1e10 + 3.14159)
        _ = v23
    }
    switch (("a" + ("\n\t" + "")) + "a") {
    case "foo bar":
        v2 = v2
    }
    v6 = v6
    var v24 uint8 = 1
    v25 := GenS2{F0: 2}
    var v26 int32 = 2147483647
    var v27 string = ("a" + "a")
    if v2 {
        switch -9223372036854775808 {
        case -266:
            _ = len(v4)
        default:
            var v28 uint16 = 255
            _ = v28
        }
        v6 = (!false)
    }
    switch "hello" {
    case "":
        fmt.Println(v26, v27)
    default:
        v27 = ""
    }
    var v29 float32 = 0.0
    fmt.Println(v27)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

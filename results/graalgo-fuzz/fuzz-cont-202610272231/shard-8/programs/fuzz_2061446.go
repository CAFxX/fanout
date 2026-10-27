package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int32, p1 int64) bool {
    p0 = (p0 & p0)
    v5 := GenS1{F0: 1000}
    v6 := []int{32767, 32767, -266}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (!false)
}

func main() {
    var v1 uint32 = 962
    var v2 uint16 = 703
    v3 := make(map[string]bool)
    var v4 uint32 = (v1 | (^(v1)))
    fmt.Println(v1, v4)
    if (!true) {
        for fi7 := 0; fi7 < 10; fi7++ {
            v3["\n\t"] = true
            _ = len(v3)
            _ = fi7
        }
        v1 += (v1 << ((0) & 63))
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS1{F0: 654}
        _ = fi10
        _ = v11
    }
    v2 -= v2
    var v12 uint8 = 127
    for fi13 := 0; fi13 < 14; fi13++ {
        switch ("hello" + "") {
        case "":
            fmt.Println(v2, v12)
        case "a":
            var v14 int32 = 1
            _ = v14
        case "\n\t":
            for fi15 := 0; fi15 < 17; fi15++ {
                if false {
                    switch (fi13 - (fi13 << ((4294967295) & 63))) {
                    case -171:
                        v2 = v2
                    case -2147483648:
                        var v16 float32 = 1.5
                        _ = v16
                    }
                    if true {
                        if true {
                            var v17 int16 = 32767
                            var v18 uint32 = 831
                            var v19 bool = ((v17 >> ((2) & 63)) != (^(v17)))
                            _ = v17
                            _ = v18
                            _ = v19
                        } else {
                            v20 := []int{(^(fi13)), 32767, (^(fi15)), -1}
                            _ = v20
                        }
                        var v21 bool = (215 >= 702)
                        var v22 uint32 = v1
                        _ = v21
                        _ = v22
                    }
                    fmt.Println(v4, v12)
                    fmt.Println(fi15)
                }
                switch (("\n\t" + "\n\t") + "a") {
                case "foo bar":
                    fmt.Println(fi13)
                case "\n\t":
                    var v23 int8 = 100
                    _ = v23
                case "a":
                    fmt.Println(fi15, v4, fi15)
                default:
                    fmt.Println(v4)
                }
                _ = fi15
            }
        default:
            switch ("a" + "\n\t") {
            case "hello":
                var v24 uint8 = 2
                _ = v24
            case "a":
                var v25 uint64 = 4294967295
                _ = v25
            default:
                v26 := GenS1{F0: -32768}
                _ = v26
            }
        }
        v27 := v3[("" + ("foo bar" + "hello"))]
        _ = fi13
        _ = v27
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        fmt.Println(fi28)
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}

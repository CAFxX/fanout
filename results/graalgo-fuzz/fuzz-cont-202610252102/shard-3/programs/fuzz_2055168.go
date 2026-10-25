package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint16 = 65535
    var v2 int = -1
    var v3 float64 = (1.5 - 1e10)
    var v4 string = "foo bar"
    v1 = 65535
    var v5 uint16 = v1
    v6 := []int{(v2 & -609), (v2 + 32767)}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 uint8 = 1
    v2 = 1
    if (v1 >= 32767) {
        var v8 int64 = -378
        var v9 uint16 = (v5 % ((0) | 1))
        v10 := GenS1{F0: (v8 >> ((131) & 63))}
        _ = v10
        _ = v8
        _ = v9
    } else {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 float64 = 1.5
            _ = fi11
            _ = v12
        }
        if true {
            fmt.Println(v2, v7, v3)
            for fi13 := 0; fi13 < 4; fi13++ {
                switch "hello" {
                case "a":
                    var v14 int8 = 100
                    _ = v14
                }
                v4 = v4
                v6 = append(v6, 2147483647)
                _ = fi13
            }
            fmt.Println(v3, v5)
        } else {
            var v15 int64 = 0
            v15 = (v15 ^ 42)
            _ = v15
        }
        v1 -= v1
    }
    if (false || false) {
        var v16 bool = ((v7 - 128) > (^(v7)))
        _ = v16
    } else {
        if (!(!(true || false))) {
            v5 += (v1 / (((^(v1))) | 1))
        } else {
            v4 = (v4 + (v4 + (v4 + v4)))
            if false {
                v4 = v4
                if len(v6) > 0 {
                    v6[0] = (v2 & 9223372036854775807)
                }
            } else {
                v17 := GenS1{F0: 9223372036854775807}
                if ((255 > (^(v5))) && true) {
                    fmt.Println(v3, v2, v1)
                    for fi18 := 0; fi18 < 18; fi18++ {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v19 bool = true
                        _ = fi18
                        _ = v19
                    }
                    v6 = append(v6, (v2 << ((65535) & 63)))
                } else {
                    _ = len(v6)
                    v3 = 1.5
                }
                var v20 uint = 42
                _ = v17
                _ = v20
            }
        }
        fmt.Println(v3, v1, v2)
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            fmt.Println(fi21)
            v7 += (v7 ^ v7)
            _ = fi22
        }
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 string = v4
            _ = fi23
            _ = v24
        }
        var v25 uint = 1
        _ = fi21
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

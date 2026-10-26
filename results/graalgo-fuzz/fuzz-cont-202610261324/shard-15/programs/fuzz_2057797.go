package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[string]int)
    var v2 int = 127
    var v3 uint8 = 144
    v3 = (v3 | v3)
    v3 -= v3
    for fi4 := 0; fi4 < 16; fi4++ {
        v2 += (-(fi4))
        v5 := []int{fi4}
        v2 -= v2
        _ = fi4
        _ = v5
    }
    if (false || true) {
        v6 := GenS1{F0: (v2 / (((^(v2))) | 1))}
        var v7 int16 = 77
        v8 := []int{v2, 0, 32767}
        _ = v6
        _ = v7
        _ = v8
    } else {
        if true {
            v9 := []int{}
            _ = v9
        } else {
            if (63 != 1) {
                v10 := v1[""]
                if ("" < "foo bar") {
                    _ = len(v1)
                }
                for fi11 := 0; fi11 < 9; fi11++ {
                    if false {
                        fmt.Println(fi11)
                        v12 := GenS1{F0: v10}
                        if (!false) {
                            var v13 int16 = 1000
                            fmt.Println(v3, v13, v2)
                            _ = v13
                        }
                        _ = v12
                    } else {
                        v2 = v2
                    }
                    switch -2147483648 {
                    case 4294967295:
                        fmt.Println(v2, v3)
                    }
                    _ = fi11
                }
                _ = v10
            } else {
                var v14 int32 = 404
                v14 += (-(v14))
                v15 := []int{(v2 - v2), (v2 / ((4294967295) | 1))}
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 3; fi16++ {
                v17 := GenS1{F0: (^(v2))}
                var v18 uint16 = 65535
                _ = fi16
                _ = v17
                _ = v18
            }
            v3 = (v3 >> ((0) & 63))
        }
        v3 += v3
        v3 = v3
    }
    v19 := []int{}
    if len(v19) > 0 {
        v19[0] = v2
    }
    _ = len(v1)
    if false {
        v2 += v2
        _ = len(v19)
        fmt.Println(v3, v3)
    }
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    var v20 int8 = 1
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[string]int)
    var v2 int64 = -1
    v3 := make(map[int]bool)
    var v4 float32 = 1.5
    v5 := make(map[string]int)
    var v6 uint32 = 42
    v7 := GenS1{F0: 2147483647}
    var v8 int64 = (v2 & (v2 ^ 626))
    for fi9 := 0; fi9 < 6; fi9++ {
        if true {
            v10 := []int{(fi9 & (fi9 / ((fi9) | 1)))}
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 uint8 = 2
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        var v13 string = "\n\t"
        v14 := v1[v13]
        _ = fi9
        _ = v13
        _ = v14
    }
    v4 += (v4 + (v4 * v4))
    for fi15 := 0; fi15 < 9; fi15++ {
        v4 = v4
        v2 -= v8
        var v16 string = ""
        _ = fi15
        _ = v16
    }
    v6 -= 2147483648
    var v17 uint16 = 42
    if false {
        if true {
            for fi18 := 0; fi18 < 5; fi18++ {
                v19 := GenS1{F0: (fi18 % ((447) | 1))}
                v20 := GenS1{F0: (-(fi18))}
                _ = fi18
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 13; fi21++ {
                var v22 uint32 = 65535
                v23 := GenS1{F0: fi21}
                v6 = v6
                _ = fi21
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 6; fi24++ {
                fmt.Println(v6, fi24, v4)
                if ("\n\t" < "\n\t") {
                    v25 := []int{(-(fi24)), (fi24 ^ fi24)}
                    _ = v25
                }
                v4 = (1e10 * -2.25)
                _ = fi24
            }
        } else {
            v26 := []int{32768, 32768}
            fmt.Println(v17, v17, v17)
            v27 := GenS1{F0: 448}
            _ = v26
            _ = v27
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

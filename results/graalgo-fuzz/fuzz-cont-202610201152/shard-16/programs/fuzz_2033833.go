package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 string }
func genFn0(p0 uint) string {
    if false {
        var v7 uint32 = 42
        for fi8 := 0; fi8 < 6; fi8++ {
            v7 -= v7
            v7 += 4294967295
            v7 = (^(v7))
            _ = fi8
        }
        _ = v7
    } else {
        if true {
            var v9 int = 32767
            v10 := []int{(v9 - v9), (v9 / (((v9 << ((p0) & 63))) | 1)), v9, 32767}
            for fi11 := 0; fi11 < 6; fi11++ {
                for fi12 := 0; fi12 < 3; fi12++ {
                    v13 := GenS1{F0: 2}
                    var v14 float64 = 0.1
                    if false {
                        p0 = 1
                        switch ("hello" + "\n\t") {
                        case "a":
                            fmt.Println(v14, fi11, p0)
                        default:
                            var v15 uint8 = 127
                            _ = v15
                        }
                        fmt.Println(p0, v14, v14)
                    } else {
                        p0 = 4294967295
                        var v16 float32 = (float32(v14) + -2.25)
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        _ = v16
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = fi11
            }
            _ = v10
            _ = v9
        }
        var v17 uint16 = 2
        _ = v17
    }
    if false {
        p0 -= p0
        p0 = p0
        for fi18 := 0; fi18 < 2; fi18++ {
            v19 := GenS1{F0: 0}
            v20 := []int{fi18, 2147483647, fi18, (fi18 << (((p0 & (p0 ^ (p0 + (p0 << (((p0 << ((0) & 63))) & 63)))))) & 63))}
            for fi21 := 0; fi21 < 5; fi21++ {
                var v22 int64 = 150
                _ = fi21
                _ = v22
            }
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    v23 := []int{}
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 uint32 = 557
        for fi26 := 0; fi26 < 12; fi26++ {
            fmt.Println(fi24, fi26)
            if true {
                var v27 uint32 = (^(v25))
                var v28 uint8 = 2
                _ = v27
                _ = v28
            }
            var v29 float64 = float64(float32(1e10))
            _ = fi26
            _ = v29
        }
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = v23
    return "foo bar"
}

func main() {
    var v1 int64 = 42
    var v2 int32 = -406
    var v3 float64 = 1e10
    v4 := make(map[int]string)
    v5 := make(map[int]int)
    var v6 int = -886
    for fi30 := 0; fi30 < 6; fi30++ {
        v31 := GenS1{F0: 53}
        var v32 uint = 838
        if (!true) {
            v33 := []int{(fi30 ^ fi30), (fi30 ^ -522), fi30}
            v1 = v1
            _ = v33
        }
        _ = fi30
        _ = v31
        _ = v32
    }
    v2 += (^(v2))
    v34 := GenS2{F0: "hello", F1: "a"}
    var v35 uint32 = 0
    v2 = v2
    switch (v6 + (v6 / ((v6) | 1))) {
    case 127:
        v36 := GenS2{F0: "", F1: "a"}
        _ = v36
    }
    v35 += (^(v35))
    v1 += (v1 & 701)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint64 }
func genFn0(p0 uint32, p1 uint) int32 {
    v5 := []int{int(p0), 32767, 1, 32768}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        switch (fi6 + fi6) {
        case 32767:
            var v7 int8 = 0
            _ = v7
        case -443:
            var v8 int64 = 835
            _ = v8
        }
        p1 -= p1
        _ = fi6
    }
    var v9 int32 = 229
    p1 = p1
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 2147483647
}

func main() {
    var v1 float64 = float64(3.14159)
    var v2 int32 = -420
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS1{F0: 65535, F1: 440, F2: 0}
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := []int{4294967295, fi10, fi10, 128}
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            var v14 int16 = 1000
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := []int{(fi10 | fi10), (^(fi10))}
        _ = fi10
        _ = v11
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v1 -= (1e10 / 0.1)
        var v17 float32 = float32(v1)
        fmt.Println(v1)
        _ = fi16
        _ = v17
    }
    var v18 string = "hello"
    for fi19 := 0; fi19 < 13; fi19++ {
        v2 = (-(v2))
        _ = len(v3)
        switch 127 {
        case 0:
            v20 := v3[v18]
            _ = v20
        }
        _ = fi19
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 uint16 = 32767
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 int32 = genFn0(2, 65535)
        for fi25 := 0; fi25 < 10; fi25++ {
            v26 := GenS1{F0: 42, F1: 2147483647, F2: 747}
            v27 := GenS1{F0: 2147483648, F1: 4294967295, F2: 2}
            _ = fi25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 20; fi28++ {
            var v29 int16 = 1000
            switch "a" {
            case "\n\t":
                if (false && false) {
                    v24 = (v2 % ((v2) | 1))
                    v30 := []int{fi28}
                    if len(v30) > 0 {
                        v30[0] = (^(fi23))
                    }
                    v29 += v29
                    _ = v30
                }
            case "hello":
                var v31 uint64 = 980
                _ = v31
            case "foo bar":
                fmt.Println(v1)
            }
            v29 = -740
            var v32 uint64 = 418
            _ = fi28
            _ = v29
            _ = v32
        }
        _ = fi23
        _ = v24
    }
    v33 := []int{32767, -170}
    v2 += v2
    var v34 uint32 = 2
    _ = len(v3)
    var v35 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 string; F2 uint32 }
type GenS3 struct { F0 string }
func genFn0() uint8 {
    v5 := []int{}
    var v6 string = "foo bar"
    fmt.Println(v6)
    _ = v5
    _ = v6
    return 42
}
func genFn1() int32 {
    var v7 int32 = 1
    v8 := GenS1{F0: -86}
    switch 1 {
    case 2147483647:
        v9 := []int{1}
        _ = v9
    }
    var v10 string = "\n\t"
    var v11 bool = (!false)
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    return -399
}

func main() {
    var v1 int64 = 42
    var v2 uint = 1
    var v3 float32 = 3.14159
    v4 := make(map[string]int)
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := []int{(fi12 & (fi12 * fi12)), fi12, (-(fi12))}
        _ = fi12
        _ = v13
    }
    fmt.Println(v1, v1, v1)
    v14 := []int{42}
    fmt.Println(v2, v3, v2)
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS2{F0: (fi15 * fi15), F1: "foo bar", F2: 803}
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 uint = 248
            _ = fi17
            _ = v18
        }
        v3 += (1e10 - v3)
        _ = fi15
        _ = v16
    }
    if (!false) {
        if false {
            for fi19 := 0; fi19 < 2; fi19++ {
                if len(v14) > 0 {
                    v14[0] = fi19
                }
                var v20 uint = 1
                _ = fi19
                _ = v20
            }
            var v21 float32 = 0.0
            for fi22 := 0; fi22 < 8; fi22++ {
                v2 += (^(v2))
                _ = fi22
            }
            _ = v21
        } else {
            for fi23 := 0; fi23 < 8; fi23++ {
                v2 = v2
                _ = fi23
            }
            if (!false) {
                _ = len(v14)
            } else {
                var v24 uint32 = 2147483648
                _ = v24
            }
        }
    }
    fmt.Println(v2, v2, v1)
    v25 := GenS1{F0: 100}
    v26 := GenS3{F0: "hello"}
    for fi27 := 0; fi27 < 1; fi27++ {
        v28 := GenS3{F0: ""}
        if false {
            switch fi27 {
            case 127:
                var v29 bool = true
                _ = v29
            case 32767:
                v3 = v3
            case 42:
                v1 = 1
            }
            v1 -= -604
            fmt.Println(v1, v3, fi27)
        }
        _ = fi27
        _ = v28
    }
    if true {
        if (genFn1() != 2147483647) {
            for fi30 := 0; fi30 < 16; fi30++ {
                if len(v14) > 0 {
                    v14[0] = fi30
                }
                var v31 uint64 = 0
                for fi32 := 0; fi32 < 9; fi32++ {
                    var v33 uint = v2
                    _ = fi32
                    _ = v33
                }
                _ = fi30
                _ = v31
            }
        } else {
            v2 = v2
            v2 -= (^(v2))
            var v34 float32 = v3
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}

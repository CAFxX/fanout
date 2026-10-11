package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0(p0 int, p1 string) uint64 {
    if false {
        for fi6 := 0; fi6 < 6; fi6++ {
            switch (p1 + "a") {
            case "":
                v7 := []int{p0, fi6, p0, -1}
                _ = v7
            case "\n\t":
                for fi8 := 0; fi8 < 4; fi8++ {
                    var v9 uint8 = 2
                    v10 := GenS1{F0: 0}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
            v11 := []int{fi6, (p0 % ((p0) | 1)), fi6}
            _ = fi6
            _ = v11
        }
        if true {
            var v12 bool = true
            _ = v12
        } else {
            if true {
                if false {
                    fmt.Println(p1, p0)
                    if (!false) {
                        v13 := []int{(-(p0))}
                        var v14 int = 32767
                        if true {
                            p0 -= v14
                        }
                        _ = v13
                        _ = v14
                    }
                }
                var v15 int64 = 0
                v15 += v15
                _ = v15
            }
        }
        p1 = "hello"
    }
    p1 = "foo bar"
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[string]int)
    var v3 int = 32767
    var v4 string = "foo bar"
    var v5 string = v4
    delete(v2, v4)
    var v16 int16 = 42
    var v17 uint32 = 131
    if (!(-1 <= 1)) {
        fmt.Println(v17, v3)
        if false {
            v18 := GenS2{F0: 111, F1: (-2.25 <= 3.14159)}
            var v19 uint = 65535
            v5 = "foo bar"
            _ = v18
            _ = v19
        } else {
            var v20 int32 = -419
            v21 := []int{}
            _ = v20
            _ = v21
        }
        v22 := GenS1{F0: -1}
        _ = v22
    }
    if false {
        v23 := GenS3{F0: -128, F1: uint8(v1)}
        _ = v23
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        _ = len(v2)
        _ = fi24
    }
    v4 = "a"
    var v25 uint8 = 42
    v26 := GenS1{F0: -1}
    if (-959 > 42) {
        var v27 bool = (1 == -9223372036854775808)
        v17 += v17
        var v28 uint32 = (v1 ^ v17)
        _ = v27
        _ = v28
    } else {
        v16 -= (v16 / ((-1) | 1))
        v29 := GenS3{F0: -93, F1: 255}
        if (v25 >= v25) {
            var v30 int8 = 100
            _ = v30
        }
        _ = v29
    }
    if (42 >= -153) {
        if (1e10 != (3.14159 / 1.5)) {
            if ((v16 & 768) != (v16 * -591)) {
                var v31 uint32 = v1
                _ = v31
            }
            var v32 string = (("foo bar" + "a") + v5)
            _ = v32
        }
    } else {
        for fi33 := 0; fi33 < 8; fi33++ {
            if (-1 == -2147483648) {
                fmt.Println(v4, v4)
                v34 := []int{(^(fi33)), v3, (v3 - (^(fi33))), (fi33 - 1)}
                _ = v34
            }
            _ = fi33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}

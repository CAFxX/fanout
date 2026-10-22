package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint32 }
func genFn0(p0 int16, p1 int32, p2 int64) int64 {
    p0 += p0
    p1 -= 1
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 string = "a"
        v7 := []int{fi5, -1, (fi5 + fi5), 42}
        switch (v6 + v6) {
        case "hello":
            for fi8 := 0; fi8 < 8; fi8++ {
                if ((((226 >= 0) && (127 < 1)) && true) || ((fi8 << ((144) & 63)) > fi5)) {
                    var v9 uint64 = 4294967295
                    switch ("a" + v6) {
                    case "hello":
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    case "\n\t":
                        fmt.Println(p0, v6, v9)
                    }
                    _ = v9
                }
                if len(v7) > 0 {
                    v7[0] = fi5
                }
                v10 := GenS1{F0: -117}
                var v11 float64 = 0.1
                _ = fi8
                _ = v10
                _ = v11
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    p2 += p2
    var v12 uint16 = 314
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return p2
}
func genFn1(p0 uint16) uint64 {
    var v13 string = "a"
    p0 += p0
    _ = p0
    _ = v13
    return 301
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{42, 1}
    var v3 int16 = 1000
    v4 := []int{4294967295, 42, 0, 2147483647}
    switch -878 {
    case -601:
        _ = len(v4)
    case 165:
        v14 := GenS1{F0: 42}
        _ = v14
    case -9223372036854775808:
        if false {
            v15 := GenS1{F0: -2147483648}
            _ = v15
        }
    }
    v3 += 914
    for fi16 := 0; fi16 < 16; fi16++ {
        fmt.Println(fi16, fi16)
        fmt.Println(v3, fi16)
        _ = fi16
    }
    if true {
        v17 := GenS2{F0: (!(false || true))}
        _ = v17
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3)
    var v18 int = -2147483648
    switch ("" + "hello") {
    case "":
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := GenS2{F0: false}
            v3 = v3
            fmt.Println(v18, v3)
            _ = fi19
            _ = v20
        }
    case "foo bar":
        v3 += 1
    case "a":
        v21 := GenS3{F0: 1, F1: 4294967295}
        _ = v21
    }
    _ = len(v4)
    v18 = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}

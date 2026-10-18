package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int }
func genFn0(p0 uint32, p1 string) int32 {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 uint64 = 18446744073709551615
        switch fi5 {
        case 4294967295:
            p1 = ("foo bar" + p1)
        case 32768:
            v7 := []int{fi5, (fi5 << ((610) & 63))}
            _ = v7
        case 42:
            p0 = p0
        }
        var v8 int8 = 100
        _ = fi5
        _ = v6
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        p0 -= p0
        _ = fi9
    }
    if (0.1 < (1.5 * -2.25)) {
        p1 = p1
    } else {
        v10 := []int{1, 2147483647, -681}
        if (!true) {
            if ((128 >= 0) || false) {
                for fi11 := 0; fi11 < 1; fi11++ {
                    switch (p1 + p1) {
                    case "":
                        _ = len(v10)
                    case "hello":
                        p1 = "a"
                    }
                    _ = fi11
                }
            } else {
                v12 := GenS1{F0: 100}
                var v13 uint32 = 2
                v14 := GenS1{F0: 1}
                _ = v12
                _ = v13
                _ = v14
            }
        } else {
            switch 42 {
            case 4294967295:
                if true {
                    p0 = 0
                    for fi15 := 0; fi15 < 12; fi15++ {
                        v16 := GenS2{F0: 72, F1: 0}
                        v17 := GenS1{F0: 100}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    if (238 > 18446744073709551615) {
                        var v18 uint8 = 87
                        if true {
                            v19 := GenS3{F0: -2147483648}
                            _ = v19
                        }
                        _ = v18
                    }
                }
            case 0:
                var v20 uint16 = 32767
                _ = v20
            }
        }
        _ = len(v10)
        _ = v10
    }
    var v21 uint16 = 1
    for fi22 := 0; fi22 < 7; fi22++ {
        v21 = 2
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = v21
    return -2147483648
}

func main() {
    var v1 uint32 = 822
    v2 := make(map[string]string)
    v3 := make(map[uint32]uint32)
    var v4 string = "hello"
    var v23 float64 = 3.14159
    if (0 < 4294967295) {
        for fi24 := 0; fi24 < 10; fi24++ {
            v23 = v23
            _ = fi24
        }
        v2[v4] = "foo bar"
        if true {
            var v25 int8 = -1
            _ = v25
        }
    } else {
        fmt.Println(v1)
        v26 := v3[v1]
        v27 := GenS2{F0: 17, F1: 255}
        _ = v26
        _ = v27
    }
    fmt.Println(v4, v1, v4)
    v1 = (^(v1))
    for fi28 := 0; fi28 < 18; fi28++ {
        var v29 int8 = 40
        v30 := GenS2{F0: v29, F1: 1}
        var v31 float64 = (v23 + v23)
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    var v32 float64 = 1.5
    var v33 uint16 = 255
    fmt.Println(v1, v1)
    fmt.Println(v4, v23)
    fmt.Println(v23, v1, v23)
    v33 += 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}

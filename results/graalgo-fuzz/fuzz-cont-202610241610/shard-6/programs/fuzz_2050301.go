package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint) uint8 {
    p0 = (p0 & (p0 % ((4294967295) | 1)))
    switch "a" {
    case "foo bar":
        if (-1 > -98) {
            var v6 uint = (p0 % ((p0) | 1))
            _ = v6
        } else {
            if (false && (128 > 1)) {
                v7 := []int{344, -242}
                p0 -= 0
                _ = v7
            } else {
                v8 := GenS1{F0: 42}
                v9 := []int{}
                switch (("hello" + "hello") + "") {
                case "a":
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
                _ = v8
                _ = v9
            }
            p0 -= (p0 * (^(p0)))
        }
    case "":
        v10 := []int{-2147483648, 2147483647, 128}
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 int = fi11
        p0 += p0
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS2{F0: true}
            v15 := GenS1{F0: -9223372036854775808}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    v16 := []int{}
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    _ = p0
    _ = v16
    return 42
}

func main() {
    var v1 int16 = 111
    v2 := make(map[uint32]bool)
    var v3 uint = 255
    v4 := make(map[int]uint32)
    var v5 bool = true
    v4[-9223372036854775808] = 672
    switch ("hello" + "a") {
    case "\n\t":
        if v5 {
            v3 += 1
            for fi17 := 0; fi17 < 11; fi17++ {
                if (!(!v5)) {
                    switch "hello" {
                    case "foo bar":
                        switch (("foo bar" + ((("hello" + "") + "foo bar") + "a")) + "hello") {
                        case "hello":
                            v18 := v2[335]
                            _ = v18
                        case "a":
                            switch "hello" {
                            case "hello":
                                var v19 uint8 = genFn0(v3)
                                _ = v19
                            default:
                                v20 := GenS1{F0: -1}
                                _ = v20
                            }
                        case "\n\t":
                            v3 += v3
                        default:
                            v21 := GenS3{F0: 2147483648}
                            _ = v21
                        }
                    default:
                        fmt.Println(v3)
                    }
                    var v22 float64 = 3.14159
                    v1 += 32767
                    _ = v22
                }
                switch 487 {
                case 128:
                    var v23 uint64 = 18446744073709551615
                    _ = v23
                default:
                    switch fi17 {
                    case 0:
                        var v24 uint8 = 127
                        _ = v24
                    default:
                        fmt.Println(v5, v1, fi17)
                    }
                }
                var v25 string = ""
                _ = fi17
                _ = v25
            }
            var v26 uint64 = 1
            _ = v26
        }
    case "foo bar":
        for fi27 := 0; fi27 < 19; fi27++ {
            var v28 float32 = float32(0.0)
            var v29 string = "a"
            _ = fi27
            _ = v28
            _ = v29
        }
    default:
        v30 := []int{}
        _ = v30
    }
    fmt.Println(v5, v5, v5)
    for fi31 := 0; fi31 < 20; fi31++ {
        var v32 int16 = -1
        for fi33 := 0; fi33 < 17; fi33++ {
            v5 = true
            v1 = (v32 + v1)
            _ = fi33
        }
        var v34 int32 = 2147483647
        _ = fi31
        _ = v32
        _ = v34
    }
    var v35 int = -1
    for fi36 := 0; fi36 < 20; fi36++ {
        fmt.Println(fi36, v1, v3)
        _ = fi36
    }
    v37 := []int{v35, v35, 4294967295}
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 int64 = 42
        v40 := GenS3{F0: 2147483647}
        for fi41 := 0; fi41 < 9; fi41++ {
            v42 := GenS1{F0: (^(v39))}
            var v43 uint32 = 2
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = fi38
        _ = v39
        _ = v40
    }
    var v44 uint = v3
    v35 += 32767
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v44
    _ = v5
}

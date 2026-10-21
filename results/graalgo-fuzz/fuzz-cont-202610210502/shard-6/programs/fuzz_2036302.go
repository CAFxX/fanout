package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 int64 = -1
    var v2 int32 = 2147483647
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    var v5 float64 = 1e10
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: 42}
    fmt.Println(v2, v2, v1)
    v8 := GenS2{F0: 100}
    v9 := GenS2{F0: 42}
    if true {
        var v10 uint64 = 940
        switch "foo bar" {
        case "\n\t":
            v2 -= v2
        case "":
            v10 = 0
        default:
            for fi11 := 0; fi11 < 6; fi11++ {
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 int16 = -32768
                    switch "a" {
                    case "":
                        v14 := GenS3{F0: v1}
                        _ = v14
                    default:
                        fmt.Println(v10, fi11, fi11)
                    }
                    _ = fi12
                    _ = v13
                }
                _ = fi11
            }
        }
        _ = v10
    } else {
        if true {
            var v15 uint16 = 42
            v16 := GenS4{F0: 435}
            _ = v15
            _ = v16
        } else {
            if (v5 > v5) {
                var v17 uint64 = 4294967295
                for fi18 := 0; fi18 < 3; fi18++ {
                    switch "a" {
                    case "foo bar":
                        if (1 == -925) {
                            for fi19 := 0; fi19 < 15; fi19++ {
                                var v20 uint64 = v17
                                var v21 uint16 = 32767
                                switch "hello" {
                                case "a":
                                    fmt.Println(v1)
                                case "hello":
                                    var v22 uint64 = v20
                                    _ = v22
                                case "foo bar":
                                    v23 := GenS2{F0: -30}
                                    _ = v23
                                default:
                                    v24 := GenS5{F0: 1000}
                                    _ = v24
                                }
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            for fi25 := 0; fi25 < 6; fi25++ {
                                var v26 int16 = 1000
                                v26 = (v26 << ((18) & 63))
                                _ = fi25
                                _ = v26
                            }
                            v1 = v1
                        } else {
                            fmt.Println(v1)
                        }
                    }
                    v27 := []int{-1, 127, fi18, fi18}
                    switch "\n\t" {
                    case "a":
                        if false {
                            _ = len(v27)
                            if (32767 < 255) {
                                var v28 bool = (v1 < (v1 >> ((439) & 63)))
                                _ = v28
                            }
                            v29 := GenS5{F0: 32767}
                            _ = v29
                        }
                    default:
                        var v30 int = 128
                        _ = v30
                    }
                    _ = fi18
                    _ = v27
                }
                v1 = 1
                _ = v17
            }
            v1 -= (v1 + (v1 + v1))
        }
        fmt.Println(v2, v1)
    }
    v5 -= v5
    v31 := v3[2147483647]
    fmt.Println(v5, v1)
    var v32 int16 = 1
    if v31 {
        var v33 bool = (-548 <= -1)
        if (true || false) {
            v2 += 2147483647
            v34 := v3[2]
            _ = v34
        }
        _ = v33
    } else {
        v32 -= 0
        v2 -= v2
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        var v36 uint16 = 255
        v2 -= (v2 ^ 42)
        switch (("a" + "") + "") {
        case "a":
            v37 := []int{}
            _ = v37
        }
        _ = fi35
        _ = v36
    }
    var v38 int32 = 2147483647
    var v39 int32 = (v2 + -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

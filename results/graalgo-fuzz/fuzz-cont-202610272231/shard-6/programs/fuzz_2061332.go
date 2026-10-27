package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int; F1 int16 }
func genFn0(p0 int8, p1 uint) uint8 {
    var v5 int64 = -795
    var v6 bool = true
    var v7 bool = true
    v6 = (!true)
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int16, p1 string) int8 {
    v9 := []int{324}
    if true {
        var v10 uint64 = 18446744073709551615
        for fi11 := 0; fi11 < 2; fi11++ {
            if false {
                var v12 int64 = 42
                v10 = v10
                _ = v12
            }
            var v13 int = fi11
            _ = fi11
            _ = v13
        }
        _ = v10
    }
    var v14 uint64 = 2
    if false {
        v15 := GenS1{F0: p1}
        if true {
            v14 += v14
            if true {
                var v16 uint16 = 42
                _ = v16
            }
        }
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v14
    _ = v9
    return 0
}

func main() {
    var v1 int8 = 21
    var v2 uint32 = 2
    v3 := make(map[uint32]bool)
    var v4 uint16 = 418
    if true {
        v4 = (v4 >> ((1) & 63))
    } else {
        fmt.Println(v4, v2)
    }
    v4 = (v4 - (v4 << ((65535) & 63)))
    var v17 float32 = (3.14159 + (1.5 / 1.5))
    if (-32768 > 32767) {
        switch "hello" {
        case "":
            var v18 int64 = 0
            _ = v18
        case "a":
            var v19 uint = 0
            _ = v19
        }
    } else {
        for fi20 := 0; fi20 < 1; fi20++ {
            for fi21 := 0; fi21 < 15; fi21++ {
                if true {
                    var v22 int = (fi21 << ((79) & 63))
                    v23 := []int{32768}
                    v24 := GenS1{F0: "a"}
                    _ = v22
                    _ = v23
                    _ = v24
                }
                if true {
                    if (false && false) {
                        var v25 int32 = 2147483647
                        var v26 int32 = (v25 ^ (v25 + 715))
                        _ = v25
                        _ = v26
                    }
                } else {
                    v27 := v3[v2]
                    _ = len(v3)
                    var v28 int8 = v1
                    switch fi20 {
                    case 4294967295:
                        var v29 uint32 = v2
                        _ = v29
                    }
                    _ = v27
                    _ = v28
                }
                var v30 bool = true
                v31 := GenS1{F0: ("a" + "a")}
                _ = fi21
                _ = v30
                _ = v31
            }
            v4 = v4
            _ = len(v3)
            _ = fi20
        }
    }
    if true {
        var v32 int32 = 132
        var v33 uint8 = 255
        for fi34 := 0; fi34 < 20; fi34++ {
            v35 := GenS2{F0: 1, F1: 42, F2: v4}
            _ = fi34
            _ = v35
        }
        _ = v32
        _ = v33
    }
    v36 := GenS3{F0: 144, F1: -1}
    var v37 uint16 = v4
    v2 = 42
    v2 = v2
    if (!true) {
        v37 += v37
    } else {
        var v38 uint8 = 255
        v39 := []int{4294967295}
        _ = v38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}

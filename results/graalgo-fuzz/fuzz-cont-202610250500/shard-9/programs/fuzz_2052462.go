package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int8; F1 bool }
func genFn0(p0 bool) int {
    if (710 < -1) {
        v6 := []int{}
        switch "foo bar" {
        case "a":
            p0 = (230 >= 32767)
        case "":
            if p0 {
                var v7 int64 = -9223372036854775808
                _ = v7
            } else {
                var v8 int32 = 42
                _ = v8
            }
        case "hello":
            var v9 string = "a"
            _ = v9
        }
        _ = v6
    } else {
        p0 = p0
        var v10 uint8 = 127
        v11 := []int{4294967295, 680}
        _ = v10
        _ = v11
    }
    if false {
        fmt.Println(p0)
    }
    v12 := GenS1{F0: 42}
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 bool = p0
        if (65535 < 534) {
            v14 = false
            fmt.Println(fi13, v14, v14)
        }
        _ = fi13
        _ = v14
    }
    v15 := []int{}
    _ = p0
    _ = v12
    _ = v15
    return 87
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]int)
    v3 := []uint32{2147483648}
    v4 := make(map[string]bool)
    v5 := make(map[int]uint32)
    if (!(!false)) {
        for fi16 := 0; fi16 < 4; fi16++ {
            v17 := GenS2{F0: (v1 * v1), F1: false}
            v1 -= (v1 + v1)
            _ = fi16
            _ = v17
        }
        v1 = (v1 - (v1 % ((v1) | 1)))
    } else {
        var v18 int8 = -59
        v19 := GenS2{F0: (v1 + 128), F1: (-995 == 4294967295)}
        _ = v18
        _ = v19
    }
    var v20 int32 = -2147483648
    switch "\n\t" {
    case "hello":
        v20 += (v20 >> ((503) & 63))
    case "\n\t":
        v1 = v1
    }
    fmt.Println(v1)
    if false {
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 int = 112
            if false {
                switch (("foo bar" + (("a" + "a") + "a")) + "hello") {
                case "foo bar":
                    v23 := GenS2{F0: (v1 + (v1 | v1)), F1: ((v1 != v1) || false)}
                    _ = v23
                case "":
                    var v24 int32 = (v20 ^ v20)
                    _ = v24
                }
                fmt.Println(v22, v1)
                for fi25 := 0; fi25 < 6; fi25++ {
                    var v26 int = (fi25 + (v22 * (fi25 % (((fi25 ^ fi21)) | 1))))
                    var v27 uint16 = 350
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
            _ = fi21
            _ = v22
        }
        fmt.Println(v1)
    } else {
        for fi28 := 0; fi28 < 15; fi28++ {
            v20 += 670
            fmt.Println(fi28, v1, v1)
            _ = fi28
        }
        v29 := GenS1{F0: 131}
        _ = v29
    }
    var v30 float64 = -2.25
    for fi31 := 0; fi31 < 7; fi31++ {
        var v32 int = fi31
        for fi33 := 0; fi33 < 16; fi33++ {
            v20 += -418
            _ = fi33
        }
        var v34 int8 = -70
        _ = fi31
        _ = v32
        _ = v34
    }
    if (("" + "") > ("\n\t" + ("hello" + "a"))) {
        v30 = v30
        v35 := GenS3{F0: -101, F1: (42 > 178)}
        _ = v35
    } else {
        for fi36 := 0; fi36 < 5; fi36++ {
            v20 -= 388
            _ = fi36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
func genFn0() int16 {
    if false {
        if true {
            v7 := []int{4294967295, 9223372036854775807, 2147483647}
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := GenS1{F0: 2}
                var v10 uint64 = 552
                v10 = (^(v10))
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 int32 = 2147483647
            _ = v11
            _ = v7
        } else {
            var v12 uint16 = 255
            v12 = 32767
            if false {
                v13 := GenS2{F0: 18446744073709551615}
                fmt.Println(v12)
                _ = v13
            } else {
                fmt.Println(v12)
                v14 := GenS2{F0: 2}
                _ = v14
            }
            _ = v12
        }
    }
    var v15 int = -589
    if false {
        v16 := GenS1{F0: 201}
        v17 := []int{(v15 >> ((255) & 63))}
        _ = v16
        _ = v17
    }
    _ = v15
    return 106
}
func genFn1(p0 int32) uint {
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 bool = false
        _ = fi18
        _ = v19
    }
    p0 = (p0 << ((0) & 63))
    _ = p0
    return 10
}

func main() {
    var v1 int64 = -793
    v2 := make(map[string]int)
    var v3 string = ""
    v4 := make(map[int]bool)
    v5 := make(map[int]string)
    var v6 int8 = 117
    v1 = v1
    v3 = ""
    switch v3 {
    case "foo bar":
        v3 = (((v3 + "a") + v3) + v3)
    default:
        var v20 bool = true
        _ = v20
    }
    v21 := v4[-1]
    var v22 uint64 = 2
    fmt.Println(v22, v22, v1)
    switch 394 {
    case 888:
        switch (v3 + (v3 + "\n\t")) {
        case "foo bar":
            v23 := GenS3{F0: 1}
            _ = v23
        case "":
            v24 := []int{-1}
            _ = v24
        case "hello":
            v25 := []int{-162, int(v1)}
            _ = v25
        default:
            if v21 {
                v2[v3] = 467
                fmt.Println(v6)
            }
        }
    }
    switch (v3 + ("\n\t" + v3)) {
    case "foo bar":
        v26 := GenS4{F0: v3}
        _ = v26
    case "":
        var v27 bool = v21
        _ = v27
    case "a":
        for fi28 := 0; fi28 < 1; fi28++ {
            var v29 uint = 482
            switch "hello" {
            case "":
                var v30 float64 = 1e10
                _ = v30
            case "hello":
                if (v21 || (255 != 255)) {
                    var v31 bool = (false && ((2 == 1) && (!((fi28 < fi28) || (!false)))))
                    _ = v31
                }
            }
            _ = fi28
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint64) int {
    var v4 int16 = -1
    var v5 float32 = 0.0
    v6 := []int{-1}
    p0 = p0
    var v7 int64 = -9223372036854775808
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 128
}
func genFn1() uint16 {
    var v8 uint16 = 889
    v9 := []int{535, 2147483647, -2147483648}
    _ = v8
    _ = v9
    return (^(v8))
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    v10 := GenS1{F0: true}
    for fi11 := 0; fi11 < 12; fi11++ {
        fmt.Println(fi11)
        switch fi11 {
        case 2147483647:
            var v12 int16 = -1
            _ = v12
        case 32767:
            var v13 uint = 543
            _ = v13
        }
        _ = fi11
    }
    if (255 > 0) {
        if true {
            v3[((("foo bar" + ("\n\t" + "\n\t")) + ("a" + ("foo bar" + ""))) + (((("foo bar" + "a") + (("\n\t" + "\n\t") + "hello")) + ("a" + "")) + "hello"))] = 0
        }
        var v14 uint16 = 1
        _ = v14
    }
    var v15 int = -9223372036854775808
    for fi16 := 0; fi16 < 17; fi16++ {
        fmt.Println(fi16)
        var v17 int64 = -474
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        switch ((("foo bar" + "hello") + "\n\t") + "a") {
        case "a":
            switch (^(fi18)) {
            case 9223372036854775807:
                if false {
                    var v19 int8 = -1
                    _ = v19
                }
            }
        case "":
            fmt.Println(v15, fi18)
        case "hello":
            v20 := GenS1{F0: (fi18 >= -2147483648)}
            _ = v20
        }
        v15 = fi18
        v21 := GenS1{F0: true}
        _ = fi18
        _ = v21
    }
    v15 += v15
    if (-1 != 0) {
        for fi22 := 0; fi22 < 17; fi22++ {
            if true {
                var v23 uint16 = 65535
                var v24 int8 = 0
                v24 += v24
                _ = v23
                _ = v24
            } else {
                v15 = fi22
            }
            _ = fi22
        }
        fmt.Println(v15)
        switch (v15 | 32768) {
        case -598:
            v3[""] = (v15 >> ((2) & 63))
        case -894:
            for fi25 := 0; fi25 < 15; fi25++ {
                v15 -= (fi25 & fi25)
                if (genFn1() >= 32767) {
                    var v26 uint32 = 42
                    _ = v26
                } else {
                    fmt.Println(fi25, fi25)
                    fmt.Println(fi25, v15)
                    fmt.Println(fi25)
                }
                _ = fi25
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
}

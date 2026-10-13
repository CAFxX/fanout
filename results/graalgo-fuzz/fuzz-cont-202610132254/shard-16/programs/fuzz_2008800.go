package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
func genFn0() string {
    var v4 uint = 1
    switch ("" + "") {
    case "hello":
        if false {
            var v5 uint16 = 717
            var v6 string = "\n\t"
            _ = v5
            _ = v6
        } else {
            switch -2147483648 {
            case -1:
                v4 = 2
            default:
                var v7 int32 = -2147483648
                _ = v7
            }
            var v8 uint8 = 42
            v8 = (v8 >> ((v4) & 63))
            _ = v8
        }
    case "\n\t":
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := []int{(fi9 + 32767), 128}
            switch (fi9 / ((fi9) | 1)) {
            case -1:
                if len(v10) > 0 {
                    v10[0] = (fi9 >> ((v4) & 63))
                }
            case 32767:
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
            var v11 int32 = -375
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 int16 = 42
    _ = v12
    _ = v4
    return "a"
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]int)
    var v13 uint8 = 1
    if true {
        v13 = v13
    } else {
        fmt.Println(v13)
        v14 := GenS1{F0: 71}
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        _ = len(v1)
        switch 128 {
        case -76:
            if (1.5 != float64(1.5)) {
                var v16 uint64 = 2
                var v17 int16 = -1
                fmt.Println(fi15, v16)
                _ = v16
                _ = v17
            }
        case -1:
            var v18 uint8 = v13
            _ = v18
        case 2147483647:
            v19 := GenS1{F0: 648}
            _ = v19
        default:
            for fi20 := 0; fi20 < 8; fi20++ {
                v1[4294967295] = 1
                v21 := GenS2{F0: -1}
                var v22 int16 = 32767
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        _ = fi15
    }
    v13 += 128
    var v23 uint64 = 2
    var v24 uint8 = 1
    v25 := []int{42}
    v23 = uint64(v24)
    fmt.Println(v23, v13, v23)
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 int16 = -1
        fmt.Println(fi26)
        _ = fi26
        _ = v27
    }
    var v28 bool = true
    for fi29 := 0; fi29 < 13; fi29++ {
        var v30 int = 42
        v24 = (^(v24))
        v24 = v24
        _ = fi29
        _ = v30
    }
    fmt.Println(v23, v28)
    if false {
        delete(v2, uint32(v23))
        for fi31 := 0; fi31 < 7; fi31++ {
            if ((!v28) && (-2.25 <= (3.14159 - 3.14159))) {
                if true {
                    fmt.Println(v24, v24)
                    for fi32 := 0; fi32 < 5; fi32++ {
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                        _ = fi32
                    }
                }
                var v33 string = ("a" + "hello")
                for fi34 := 0; fi34 < 10; fi34++ {
                    switch v33 {
                    case "foo bar":
                        if len(v25) > 0 {
                            v25[0] = fi34
                        }
                    }
                    _ = fi34
                }
                _ = v33
            }
            v35 := GenS3{F0: 1}
            v24 += 42
            _ = fi31
            _ = v35
        }
        var v36 uint8 = v13
        _ = v36
    } else {
        fmt.Println(v23)
        v2[65535] = 255
    }
    fmt.Println(v23, v28, v24)
    var v37 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v37
}

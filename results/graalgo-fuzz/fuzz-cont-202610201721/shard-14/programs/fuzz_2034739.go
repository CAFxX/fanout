package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int8; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 string; F2 int }

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 127
    var v3 int = -2147483648
    v2 += v2
    switch (v3 - v3) {
    case 9223372036854775807:
        switch "\n\t" {
        case "\n\t":
            for fi4 := 0; fi4 < 18; fi4++ {
                for fi5 := 0; fi5 < 12; fi5++ {
                    v6 := GenS1{F0: int32(v2)}
                    _ = fi5
                    _ = v6
                }
                _ = fi4
            }
        case "":
            var v7 string = "hello"
            _ = v7
        case "a":
            var v8 uint = 706
            _ = v8
        }
    case -2147483648:
        if true {
            var v9 string = "a"
            _ = len(v1)
            _ = v9
        }
    default:
        switch v3 {
        case -9223372036854775808:
            fmt.Println(v2)
        case 4294967295:
            v10 := GenS1{F0: -1}
            _ = v10
        default:
            if true {
                v3 = v3
                fmt.Println(v3, v3)
            } else {
                v2 = v2
                if false {
                    switch v3 {
                    case 128:
                        v11 := GenS2{F0: 368, F1: (v2 * v2), F2: (true && true)}
                        _ = v11
                    case 9223372036854775807:
                        var v12 string = "hello"
                        _ = v12
                    }
                    var v13 string = ""
                    var v14 int64 = -1
                    _ = v13
                    _ = v14
                } else {
                    if false {
                        v15 := []int{v3, (v3 + v3), v3, (^(v3))}
                        _ = v15
                    }
                }
            }
        }
    }
    if (!(255 < 128)) {
        if (v2 >= (v2 & 42)) {
            if false {
                v16 := GenS1{F0: 42}
                _ = v16
            } else {
                v17 := GenS3{F0: 2147483648}
                _ = v17
            }
            if (255 >= 134) {
                v18 := []int{}
                v2 = (v2 & v2)
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := []int{v3, v3, 2147483647}
            switch ("a" + ("foo bar" + ("\n\t" + ""))) {
            case "foo bar":
                var v21 string = "a"
                _ = v21
            }
            _ = fi19
            _ = v20
        }
    }
    v3 = (v3 - (v3 + (v3 - v3)))
    v3 -= v3
    var v22 bool = (!true)
    v2 = (v2 ^ v2)
    for fi23 := 0; fi23 < 7; fi23++ {
        v22 = true
        switch (fi23 * v3) {
        case -2147483648:
            var v24 bool = true
            _ = v24
        case 42:
            fmt.Println(fi23, fi23, v2)
        default:
            fmt.Println(v22)
        }
        _ = fi23
    }
    v22 = v22
    for fi25 := 0; fi25 < 2; fi25++ {
        switch "a" {
        case "\n\t":
            var v26 uint64 = 2
            _ = v26
        default:
            v27 := GenS4{F0: 42, F1: "foo bar", F2: v3}
            _ = v27
        }
        var v28 int = v3
        _ = fi25
        _ = v28
    }
    var v29 bool = (1 > 32767)
    _ = len(v1)
    if v29 {
        var v30 uint32 = 42
        var v31 int32 = 1
        _ = v30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint16; F2 int64 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]string)
    var v3 int = 531
    v2[112] = "foo bar"
    for fi4 := 0; fi4 < 13; fi4++ {
        v1 = 42
        _ = fi4
    }
    switch "" {
    case "":
        for fi5 := 0; fi5 < 3; fi5++ {
            if false {
                v3 = fi5
                for fi6 := 0; fi6 < 1; fi6++ {
                    switch (fi5 * fi5) {
                    case 4294967295:
                        v7 := []int{v3}
                        _ = v7
                    case 42:
                        v1 = (v1 >> ((v1) & 63))
                    case 9223372036854775807:
                        var v8 string = "foo bar"
                        _ = v8
                    }
                    if ("" < "foo bar") {
                        fmt.Println(v1)
                        v1 = (v1 / (((v1 & v1)) | 1))
                        fmt.Println(fi5, fi6, fi6)
                    } else {
                        var v9 uint8 = 16
                        _ = v9
                    }
                    _ = fi6
                }
            }
            _ = fi5
        }
    case "\n\t":
        for fi10 := 0; fi10 < 12; fi10++ {
            if ((false || false) || (127 != -1)) {
                v3 = (fi10 | (v3 << ((v1) & 63)))
                for fi11 := 0; fi11 < 15; fi11++ {
                    v12 := v2[172]
                    v13 := GenS1{F0: 0, F1: 2147483647}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                if (!(42 > -124)) {
                    var v14 uint32 = 214
                    _ = v14
                }
            } else {
                var v15 string = ("hello" + "hello")
                v16 := GenS1{F0: 42, F1: int32(v3)}
                _ = v15
                _ = v16
            }
            v17 := []int{fi10, (v3 ^ fi10)}
            v18 := GenS2{F0: 4294967295, F1: 0}
            _ = fi10
            _ = v17
            _ = v18
        }
    case "a":
        if (-293 == -246) {
            switch (v3 / ((v3) | 1)) {
            case -2147483648:
                fmt.Println(v1, v3)
            }
            v1 = (^(v1))
            v19 := GenS1{F0: 124, F1: -1}
            _ = v19
        }
    default:
        v3 -= v3
    }
    v20 := []int{(v3 % ((127) | 1)), -876, (v3 << (((v1 | (v1 - v1))) & 63)), v3}
    delete(v2, 4294967295)
    var v21 int16 = -32768
    switch "foo bar" {
    case "\n\t":
        var v22 int = (v3 >> ((4294967295) & 63))
        _ = v22
    case "a":
        if true {
            v20 = append(v20, v3)
            fmt.Println(v3)
            var v23 uint32 = 2147483648
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        fmt.Println(v1, fi24, v3)
        _ = fi24
    }
    fmt.Println(v3)
    for fi25 := 0; fi25 < 3; fi25++ {
        v1 = v1
        v26 := GenS2{F0: 1, F1: 199}
        delete(v2, 42)
        _ = fi25
        _ = v26
    }
    switch -9223372036854775808 {
    case 42:
        v27 := GenS3{F0: "foo bar"}
        _ = v27
    case -2147483648:
        var v28 uint32 = 133
        _ = v28
    case -9223372036854775808:
        for fi29 := 0; fi29 < 19; fi29++ {
            v30 := GenS4{F0: fi29, F1: 42, F2: int64(v21)}
            switch ("hello" + "\n\t") {
            case "hello":
                if true {
                    var v31 int64 = 1
                    _ = v31
                }
            }
            v32 := GenS3{F0: ""}
            _ = fi29
            _ = v30
            _ = v32
        }
    default:
        var v33 int32 = 618
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}

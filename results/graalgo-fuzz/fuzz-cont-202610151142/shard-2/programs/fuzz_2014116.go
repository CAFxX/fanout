package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int; F2 string }
type GenS5 struct { F0 uint16; F1 uint8 }
func genFn0(p0 bool) uint64 {
    v6 := []int{566}
    var v7 int16 = 1000
    p0 = (1 < 0)
    var v8 int8 = 42
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 275
}

func main() {
    var v1 uint64 = 2
    var v2 uint64 = 9
    var v3 uint8 = 255
    var v4 uint16 = 65535
    var v5 int64 = -9223372036854775808
    switch ("hello" + "\n\t") {
    case "foo bar":
        fmt.Println(v1)
    }
    var v9 int = 32768
    v10 := GenS1{F0: "foo bar", F1: 32767, F2: 218}
    v9 += v9
    v11 := GenS1{F0: "a", F1: -190, F2: (v4 | v4)}
    var v12 bool = (!((v9 | v9) >= (v9 + v9)))
    v1 = genFn0(false)
    if v12 {
        if v12 {
            v13 := []int{}
            _ = v13
        }
    } else {
        v14 := []int{v9}
        if v12 {
            if len(v14) > 0 {
                v14[0] = (v9 ^ (v9 | 32767))
            }
            switch v9 {
            case -774:
                switch "a" {
                case "hello":
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                }
            default:
                var v15 int = (v9 >> ((255) & 63))
                _ = v15
            }
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        }
        _ = v14
    }
    v16 := []int{v9, 4294967295, v9, v9}
    for fi17 := 0; fi17 < 20; fi17++ {
        if true {
            var v18 int8 = 127
            _ = v18
        }
        var v19 uint32 = 0
        fmt.Println(v5)
        _ = fi17
        _ = v19
    }
    v4 += v4
    switch (("hello" + "a") + ("a" + (("foo bar" + "") + ("hello" + (("\n\t" + "\n\t") + ("\n\t" + ("hello" + "\n\t"))))))) {
    case "":
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 uint64 = v1
            v22 := GenS2{F0: v12}
            _ = fi20
            _ = v21
            _ = v22
        }
    case "foo bar":
        switch -9223372036854775808 {
        case -931:
            v23 := GenS2{F0: true}
            _ = v23
        case -2147483648:
            for fi24 := 0; fi24 < 16; fi24++ {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                v5 += 98
                switch v9 {
                case -21:
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                case -447:
                    for fi25 := 0; fi25 < 20; fi25++ {
                        v12 = false
                        var v26 int16 = -32768
                        switch "foo bar" {
                        case "":
                            v3 += (v3 / ((v3) | 1))
                        }
                        _ = fi25
                        _ = v26
                    }
                case 127:
                    for fi27 := 0; fi27 < 18; fi27++ {
                        if true {
                            var v28 int32 = 2147483647
                            _ = v28
                        }
                        _ = fi27
                    }
                default:
                    var v29 int16 = 1
                    _ = v29
                }
                _ = fi24
            }
        case 1:
            v12 = true
        }
    case "hello":
        if v12 {
            v16 = append(v16, v9)
            for fi30 := 0; fi30 < 9; fi30++ {
                var v31 uint = 0
                v5 += (v5 & v5)
                v32 := GenS3{F0: 808}
                _ = fi30
                _ = v31
                _ = v32
            }
            var v33 uint32 = 993
            _ = v33
        } else {
            v34 := GenS4{F0: v3, F1: (v9 ^ 9223372036854775807), F2: "a"}
            switch "hello" {
            case "":
                var v35 uint = 4294967295
                _ = v35
            case "foo bar":
                v12 = false
            }
            _ = v34
        }
    default:
        if len(v16) > 0 {
            v16[0] = v9
        }
    }
    fmt.Println(v2, v1)
    _ = len(v16)
    v36 := GenS5{F0: v4, F1: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v9
}

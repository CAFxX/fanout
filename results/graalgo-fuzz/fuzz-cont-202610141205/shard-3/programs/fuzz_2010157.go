package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 uint64 }

func main() {
    var v1 int64 = 0
    var v2 string = "\n\t"
    var v3 int32 = 0
    switch ((v2 + v2) + v2) {
    case "hello":
        v4 := []int{-9223372036854775808, -323, 128}
        _ = v4
    default:
        for fi5 := 0; fi5 < 3; fi5++ {
            var v6 uint64 = 453
            for fi7 := 0; fi7 < 15; fi7++ {
                fmt.Println(fi7)
                var v8 uint16 = 42
                fmt.Println(v8)
                _ = fi7
                _ = v8
            }
            var v9 string = "a"
            _ = fi5
            _ = v6
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        if true {
            if ((v3 % (((^(v3))) | 1)) == (v3 * (v3 << ((522) & 63)))) {
                fmt.Println(v1)
                fmt.Println(fi10, fi10, v3)
                switch ("a" + "\n\t") {
                case "hello":
                    v11 := []int{}
                    _ = v11
                case "foo bar":
                    v12 := GenS1{F0: 1000, F1: v2}
                    _ = v12
                default:
                    switch "\n\t" {
                    case "\n\t":
                        for fi13 := 0; fi13 < 11; fi13++ {
                            v14 := []int{(-(fi10))}
                            if true {
                                var v15 bool = (false || false)
                                _ = v15
                            }
                            var v16 uint32 = 65535
                            _ = fi13
                            _ = v14
                            _ = v16
                        }
                    case "foo bar":
                        switch (fi10 * -198) {
                        case 42:
                            switch -1 {
                            case 32767:
                                switch "a" {
                                case "hello":
                                    for fi17 := 0; fi17 < 19; fi17++ {
                                        var v18 int = -2147483648
                                        _ = fi17
                                        _ = v18
                                    }
                                default:
                                    v19 := []int{(fi10 >> ((65535) & 63)), fi10, (fi10 | fi10)}
                                    _ = v19
                                }
                            case 1:
                                v2 = v2
                            case 4294967295:
                                v20 := GenS2{F0: (fi10 >> ((18446744073709551615) & 63))}
                                _ = v20
                            default:
                                v2 = v2
                            }
                        case -9223372036854775808:
                            v21 := []int{}
                            _ = v21
                        }
                    case "":
                        v22 := []int{-1, fi10}
                        _ = v22
                    default:
                        v1 = v1
                    }
                }
            }
            v23 := []int{(fi10 | fi10)}
            _ = v23
        }
        v24 := []int{fi10, fi10}
        if true {
            v2 = ("\n\t" + v2)
        }
        _ = fi10
        _ = v24
    }
    var v25 uint16 = 0
    fmt.Println(v3)
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 int16 = 0
        for fi28 := 0; fi28 < 9; fi28++ {
            v2 = (v2 + v2)
            if false {
                v2 = (v2 + v2)
            }
            _ = fi28
        }
        var v29 uint8 = 42
        _ = fi26
        _ = v27
        _ = v29
    }
    v30 := GenS3{F0: false, F1: 2}
    v31 := GenS1{F0: 1000, F1: ("\n\t" + v2)}
    var v32 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}

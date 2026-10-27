package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint32; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 int16) bool {
    var v5 bool = (true && false)
    var v6 uint16 = 678
    fmt.Println(v6)
    _ = p0
    _ = v5
    _ = v6
    return (1 > 1)
}

func main() {
    var v1 int64 = -1
    var v2 int = 42
    v3 := make(map[int]bool)
    var v4 uint = 33
    fmt.Println(v1, v1, v1)
    v7 := GenS1{F0: (v2 - 756), F1: (((!(false && false)) || true) && ("\n\t" > "foo bar")), F2: 783}
    var v8 float64 = 0.1
    v9 := []int{-768, (v2 >> (((v4 ^ (v4 - (v4 / ((v4) | 1))))) & 63)), v2, v2}
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(fi10, v2)
        v11 := GenS2{F0: true, F1: 511, F2: (("hello" + "hello") + ((("\n\t" + "a") + ("hello" + "\n\t")) + "foo bar"))}
        _ = fi10
        _ = v11
    }
    var v12 uint = (^(v4))
    if false {
        var v13 bool = false
        v14 := GenS3{F0: 2147483647}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        fmt.Println(v4)
        _ = fi15
    }
    fmt.Println(v2, v2)
    if (-128 == -1) {
        fmt.Println(v8, v8, v4)
        var v16 string = ""
        if false {
            v16 = ("hello" + "\n\t")
        } else {
            _ = len(v9)
            switch v16 {
            case "":
                for fi17 := 0; fi17 < 17; fi17++ {
                    for fi18 := 0; fi18 < 10; fi18++ {
                        var v19 int16 = 42
                        if true {
                            var v20 bool = false
                            _ = v20
                        } else {
                            _ = len(v9)
                            switch (v16 + v16) {
                            case "\n\t":
                                var v21 uint64 = 0
                                _ = v21
                            }
                        }
                        v3[v2] = (!(float64(float32(1.5)) < float64(3.14159)))
                        var v22 uint64 = 1
                        _ = fi18
                        _ = v19
                        _ = v22
                    }
                    v8 = 1.5
                    _ = fi17
                }
            case "hello":
                v3[v2] = true
            }
        }
        _ = v16
    }
    v23 := GenS1{F0: (v2 % ((4294967295) | 1)), F1: false, F2: 866}
    for fi24 := 0; fi24 < 6; fi24++ {
        switch "foo bar" {
        case "a":
            v4 += v12
        case "":
            switch "" {
            case "foo bar":
                for fi25 := 0; fi25 < 17; fi25++ {
                    v26 := GenS3{F0: 2}
                    _ = fi25
                    _ = v26
                }
            case "":
                if (true || (-1 < 1)) {
                    for fi27 := 0; fi27 < 5; fi27++ {
                        delete(v3, (^(fi27)))
                        if true {
                            var v28 int = fi27
                            switch fi24 {
                            case 2147483647:
                                var v29 uint8 = 0
                                _ = v29
                            }
                            _ = v28
                        }
                        v30 := GenS4{F0: "foo bar", F1: false}
                        _ = fi27
                        _ = v30
                    }
                } else {
                    for fi31 := 0; fi31 < 5; fi31++ {
                        var v32 uint64 = 42
                        var v33 float64 = (((v8 / v8) - (3.14159 - (v8 / -2.25))) - 0.0)
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                }
            case "a":
                v34 := GenS4{F0: ("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))), F1: ((!true) && true)}
                _ = v34
            default:
                var v35 uint32 = 2147483647
                _ = v35
            }
        case "\n\t":
            v2 = (fi24 % (((fi24 & fi24)) | 1))
        }
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}

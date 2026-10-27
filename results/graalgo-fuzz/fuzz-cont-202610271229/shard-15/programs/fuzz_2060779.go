package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 int8; F2 int16 }

func main() {
    var v1 uint = 0
    v2 := []string{}
    var v3 uint16 = 1
    v4 := make(map[string]uint32)
    var v5 int16 = 1000
    v6 := GenS1{F0: 423}
    v7 := GenS1{F0: 2}
    v8 := GenS2{F0: 42}
    switch int(v3) {
    case 2147483647:
        fmt.Println(v1, v3, v1)
    case 32768:
        if false {
            fmt.Println(v3)
            var v9 uint16 = (v3 << ((822) & 63))
            v5 = (v5 - 42)
            _ = v9
        }
    }
    v1 -= v1
    for fi10 := 0; fi10 < 19; fi10++ {
        v3 -= v3
        _ = fi10
    }
    v1 = v1
    v2 = append(v2, "")
    if (0 == -1) {
        v3 -= 0
        for fi11 := 0; fi11 < 2; fi11++ {
            delete(v4, ("" + "\n\t"))
            _ = fi11
        }
    } else {
        v2 = append(v2, ("" + ("hello" + "hello")))
        switch 0 {
        case -96:
            switch -1 {
            case 0:
                for fi12 := 0; fi12 < 19; fi12++ {
                    fmt.Println(fi12, v1, v3)
                    if ((v3 % (((^(v3))) | 1)) != v3) {
                        var v13 int = -9223372036854775808
                        _ = v13
                    }
                    if false {
                        if true {
                            if false {
                                v5 -= 0
                                var v14 int = 42
                                var v15 int64 = 9223372036854775807
                                _ = v14
                                _ = v15
                            } else {
                                var v16 uint = (v1 / ((v1) | 1))
                                if (695 != 62) {
                                    v5 = (v5 ^ v5)
                                    fmt.Println(v3)
                                }
                                var v17 uint16 = (^(v3))
                                _ = v16
                                _ = v17
                            }
                        }
                        var v18 int64 = 42
                        switch ("a" + ("a" + "\n\t")) {
                        case "\n\t":
                            v18 -= v18
                        case "foo bar":
                            var v19 string = ((((("a" + "hello") + ("" + "a")) + (("" + ((("hello" + "\n\t") + ("a" + "a")) + "foo bar")) + "hello")) + (((("" + ("hello" + "hello")) + "a") + (("\n\t" + ("hello" + "a")) + "")) + ("" + ((("\n\t" + "a") + "\n\t") + "\n\t")))) + "")
                            _ = v19
                        }
                        _ = v18
                    }
                    _ = fi12
                }
            case 127:
                fmt.Println(v5, v5, v3)
            case -307:
                var v20 uint8 = 42
                _ = v20
            }
        case -984:
            for fi21 := 0; fi21 < 13; fi21++ {
                if ((true || (!((!true) && (!false)))) || (2147483648 >= 4294967295)) {
                    if false {
                        v1 = v1
                        v3 = (v3 - (v3 + (^(v3))))
                        v5 = v5
                    } else {
                        v22 := GenS3{F0: 0}
                        var v23 uint = (v1 | v1)
                        if (0 == 100) {
                            switch (-(fi21)) {
                            case 128:
                                v23 = v1
                            }
                        }
                        _ = v22
                        _ = v23
                    }
                } else {
                    v1 = v1
                    for fi24 := 0; fi24 < 3; fi24++ {
                        switch 42 {
                        case -9223372036854775808:
                            var v25 int32 = -2147483648
                            switch "a" {
                            case "a":
                                v26 := GenS2{F0: -918}
                                _ = v26
                            case "\n\t":
                                var v27 string = (("foo bar" + "a") + ((("hello" + "hello") + ("hello" + "")) + (("hello" + (("\n\t" + ("foo bar" + "")) + "a")) + "foo bar")))
                                _ = v27
                            default:
                                var v28 uint32 = 673
                                _ = v28
                            }
                            _ = v25
                        case 32767:
                            var v29 int = fi21
                            _ = v29
                        default:
                            v1 = v1
                        }
                        v1 += (v1 / (((v1 << ((798) & 63))) | 1))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi24
                    }
                }
                v5 = (v5 - v5)
                _ = fi21
            }
        case -264:
            v1 += v1
        }
    }
    var v30 int32 = int32(v1)
    var v31 int = -478
    v32 := GenS4{F0: (^(v31)), F1: 127, F2: (^(v5))}
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32; F1 int32; F2 string }
type GenS3 struct { F0 int16; F1 uint16; F2 uint }
type GenS4 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    var v2 bool = true
    var v3 int = 755
    var v4 float64 = 0.0
    v5 := GenS1{F0: (("\n\t" + "") + "hello"), F1: 603}
    var v6 uint = 65535
    var v7 float64 = 1.5
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 float32 = 1e10
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: "\n\t", F1: -1}
    v1[(("" + ("foo bar" + "hello")) + "\n\t")] = true
    var v11 int64 = 639
    if v2 {
        var v12 int = (v3 / (((v3 >> ((42) & 63))) | 1))
        _ = v12
    } else {
        if (42 != 237) {
            fmt.Println(v7)
            for fi13 := 0; fi13 < 13; fi13++ {
                v14 := []int{}
                _ = fi13
                _ = v14
            }
        }
        v4 = v4
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS1{F0: "a", F1: (^(v11))}
        v17 := GenS1{F0: ("hello" + "foo bar"), F1: 0}
        var v18 float32 = -2.25
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    if (v6 >= 65535) {
        v11 += int64(v3)
        v19 := GenS2{F0: 2147483647, F1: -1, F2: ""}
        _ = v19
    } else {
        v20 := []int{v3}
        for fi21 := 0; fi21 < 20; fi21++ {
            switch (("hello" + "a") + ("" + "hello")) {
            case "":
                v22 := GenS3{F0: -32768, F1: 0, F2: v6}
                _ = v22
            case "\n\t":
                switch -1 {
                case 2147483647:
                    v23 := GenS3{F0: 32767, F1: 65535, F2: v6}
                    _ = v23
                }
            default:
                fmt.Println(v6)
            }
            _ = fi21
        }
        var v24 bool = v2
        _ = v20
        _ = v24
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        for fi26 := 0; fi26 < 10; fi26++ {
            v1[(("hello" + "\n\t") + "hello")] = (!false)
            var v27 uint32 = 249
            _ = fi26
            _ = v27
        }
        v6 = v6
        _ = fi25
    }
    if (uint32(v6) <= 846) {
        v6 = v6
        switch (((("\n\t" + (((("foo bar" + "a") + "\n\t") + ("foo bar" + "")) + "foo bar")) + (("hello" + "a") + "")) + "a") + "foo bar") {
        case "a":
            switch (v3 * v3) {
            case 1:
                v3 -= 4294967295
            case 0:
                if v2 {
                    for fi28 := 0; fi28 < 1; fi28++ {
                        fmt.Println(fi28, v7)
                        switch fi28 {
                        case -262:
                            v7 -= (float64((float32(v7) + 1e10)) * v7)
                        case -2147483648:
                            if false {
                                if v2 {
                                    if false {
                                        delete(v1, ((("\n\t" + "a") + ("" + "foo bar")) + "hello"))
                                    }
                                    v2 = false
                                    for fi29 := 0; fi29 < 3; fi29++ {
                                        v7 -= (v4 - -2.25)
                                        _ = fi29
                                    }
                                }
                            }
                        }
                        switch 231 {
                        case 572:
                            var v30 int8 = -1
                            _ = v30
                        }
                        _ = fi28
                    }
                    delete(v1, ((("a" + ((("a" + "foo bar") + "hello") + "foo bar")) + ("\n\t" + "")) + "foo bar"))
                }
            case 2147483647:
                fmt.Println(v6)
            default:
                for fi31 := 0; fi31 < 9; fi31++ {
                    v2 = (!true)
                    v2 = false
                    _ = fi31
                }
            }
        }
        v32 := []int{(v3 / (((v3 - 128)) | 1)), v3, v3, (v3 ^ -1)}
        _ = v32
    } else {
        fmt.Println(v6)
        var v33 string = ("a" + "foo bar")
        _ = v33
    }
    switch "foo bar" {
    case "hello":
        v34 := []int{}
        _ = v34
    }
    fmt.Println(v7, v7, v6)
    v3 += (v3 % (((v3 << ((1) & 63))) | 1))
    v35 := GenS4{F0: 42, F1: 100}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

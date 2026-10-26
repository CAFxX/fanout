package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint8 }

func main() {
    v1 := []string{}
    var v2 int8 = -1
    var v3 int8 = 60
    for fi4 := 0; fi4 < 17; fi4++ {
        if true {
            v5 := GenS1{F0: 65535, F1: 2, F2: 32767}
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v1 = append(v1, "hello")
                for fi6 := 0; fi6 < 4; fi6++ {
                    _ = len(v1)
                    _ = fi6
                }
            } else {
                fmt.Println(fi4, v2, v3)
            }
            _ = v5
        }
        _ = fi4
    }
    v3 += (v2 & (^(v3)))
    v2 += 1
    if len(v1) > 0 {
        v1[0] = ""
    }
    var v7 int16 = 1000
    if (false || true) {
        if (!(!((65535 > 0) && (-500 < 1)))) {
            var v8 uint16 = 65535
            v9 := GenS2{F0: ("\n\t" + ("\n\t" + "a")), F1: 128}
            _ = v8
            _ = v9
        } else {
            if (0 > 1) {
                for fi10 := 0; fi10 < 9; fi10++ {
                    var v11 string = ("foo bar" + (("a" + ("\n\t" + "")) + (("a" + "a") + "hello")))
                    for fi12 := 0; fi12 < 14; fi12++ {
                        var v13 uint = 4294967295
                        var v14 uint64 = 2
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = fi10
                    _ = v11
                }
            }
        }
        if false {
            v15 := GenS1{F0: 2147483647, F1: 1, F2: v7}
            switch 4294967295 {
            case -1:
                for fi16 := 0; fi16 < 10; fi16++ {
                    v2 = v2
                    v2 = 42
                    _ = fi16
                }
            case -987:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 128:
                switch "foo bar" {
                case "a":
                    var v17 uint = 18446744073709551615
                    _ = v17
                case "\n\t":
                    var v18 uint32 = 2147483648
                    _ = v18
                }
            }
            _ = v15
        } else {
            v19 := GenS2{F0: "\n\t", F1: 255}
            v7 += (v7 | (^(v7)))
            _ = v19
        }
    }
    var v20 int = 127
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 uint = 4294967295
        var v23 int = fi21
        var v24 float32 = 0.1
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    var v25 string = "a"
    for fi26 := 0; fi26 < 4; fi26++ {
        switch v25 {
        case "foo bar":
            v27 := GenS3{F0: 930, F1: 250}
            _ = v27
        case "a":
            v3 = -104
        default:
            v28 := GenS3{F0: 258, F1: 2}
            _ = v28
        }
        _ = fi26
    }
    v7 = (v7 * (v7 | (^(v7))))
    switch 32767 {
    case -614:
        var v29 float64 = -2.25
        _ = v29
    }
    fmt.Println(v2, v20, v3)
    v25 = v25
    switch (((v25 + v25) + ("hello" + ("" + v25))) + v25) {
    case "":
        if len(v1) > 0 {
            v1[0] = v25
        }
    }
    var v30 int8 = (^(v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v30
    _ = v7
}

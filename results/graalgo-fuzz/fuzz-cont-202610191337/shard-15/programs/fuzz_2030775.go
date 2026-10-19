package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint16; F1 int; F2 string }
type GenS4 struct { F0 int64; F1 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{(("hello" + "hello") + (((("\n\t" + "\n\t") + ("" + ("" + "\n\t"))) + ("a" + ("a" + "\n\t"))) + "\n\t")), ((("hello" + ("a" + "a")) + "foo bar") + "\n\t"), ""}
    var v3 uint = 255
    switch ("\n\t" + "") {
    case "\n\t":
        if false {
            fmt.Println(v3)
            for fi4 := 0; fi4 < 10; fi4++ {
                for fi5 := 0; fi5 < 2; fi5++ {
                    var v6 uint32 = 2
                    _ = len(v1)
                    _ = fi5
                    _ = v6
                }
                _ = fi4
            }
        } else {
            switch ("hello" + (("a" + "hello") + "foo bar")) {
            case "":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "foo bar":
                for fi7 := 0; fi7 < 10; fi7++ {
                    if true {
                        v1[42] = 428
                    }
                    _ = fi7
                }
            }
            v8 := GenS1{F0: true, F1: "", F2: v3}
            _ = v8
        }
    case "foo bar":
        v9 := GenS2{F0: 852, F1: v3}
        _ = v9
    default:
        if false {
            if ((!false) || false) {
                v10 := GenS3{F0: 1, F1: 4294967295, F2: ""}
                _ = v10
            } else {
                v3 = v3
            }
            v3 = (v3 << (((v3 - (v3 ^ 65535))) & 63))
            v2 = append(v2, "\n\t")
        } else {
            switch ("hello" + ("" + (("hello" + "foo bar") + (((((("hello" + ("a" + (("" + ((("a" + "a") + ("hello" + "foo bar")) + "\n\t")) + ("a" + "hello")))) + ("\n\t" + "hello")) + "") + "a") + "hello") + ((("\n\t" + "hello") + "hello") + ("foo bar" + "foo bar")))))) {
            case "\n\t":
                var v11 uint8 = 2
                _ = v11
            case "hello":
                _ = len(v1)
            case "foo bar":
                switch "\n\t" {
                case "":
                    var v12 uint = (v3 >> ((v3) & 63))
                    _ = v12
                case "a":
                    v3 = (v3 ^ v3)
                default:
                    var v13 string = ""
                    _ = v13
                }
            }
        }
    }
    delete(v1, 1)
    delete(v1, 376)
    v14 := GenS4{F0: 42, F1: "foo bar"}
    fmt.Println(v3)
    switch "" {
    case "\n\t":
        v2 = append(v2, "")
    case "hello":
        var v15 float32 = (-2.25 + float32(1e10))
        _ = v15
    case "foo bar":
        for fi16 := 0; fi16 < 8; fi16++ {
            for fi17 := 0; fi17 < 5; fi17++ {
                v2 = append(v2, "hello")
                v2 = append(v2, ("foo bar" + ((((("" + "") + "hello") + "hello") + "\n\t") + "a")))
                var v18 int16 = 206
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    }
    v19 := GenS4{F0: 614, F1: "\n\t"}
    var v20 int = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}

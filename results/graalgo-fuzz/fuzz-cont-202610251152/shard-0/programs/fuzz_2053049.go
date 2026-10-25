package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 2
    v3 := make(map[string]string)
    v4 := []uint32{42, 42}
    v5 := []string{}
    fmt.Println(v2)
    if false {
        fmt.Println(v2)
        v2 = 42
    } else {
        _ = len(v3)
        switch "a" {
        case "hello":
            var v6 int = 638
            _ = v6
        default:
            v7 := v1[877]
            _ = v7
        }
    }
    v2 = (^(v2))
    if (false || true) {
        v2 -= 65535
        if (true || (65535 != 2147483647)) {
            if (1 > -38) {
                switch ("\n\t" + "") {
                case "hello":
                    fmt.Println(v2)
                case "a":
                    var v8 uint8 = 127
                    _ = v8
                case "":
                    v2 += (v2 % ((1) | 1))
                }
                for fi9 := 0; fi9 < 20; fi9++ {
                    _ = len(v3)
                    switch (fi9 / (((fi9 + fi9)) | 1)) {
                    case 334:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case 0:
                        v4 = append(v4, 481)
                    }
                    v10 := v1[2147483648]
                    _ = fi9
                    _ = v10
                }
                v11 := GenS1{F0: "\n\t", F1: 1}
                _ = v11
            } else {
                if true {
                    var v12 int16 = 1000
                    _ = v12
                }
            }
            for fi13 := 0; fi13 < 15; fi13++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v2 -= (v2 / ((v2) | 1))
                _ = fi13
            }
            var v14 uint8 = 42
            _ = v14
        }
    } else {
        switch 4294967295 {
        case 9223372036854775807:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            var v15 uint16 = 868
            _ = v15
        }
        _ = len(v1)
    }
    v2 = (^(v2))
    v16 := GenS1{F0: "", F1: 100}
    for fi17 := 0; fi17 < 15; fi17++ {
        fmt.Println(fi17)
        v18 := GenS2{F0: 65535}
        for fi19 := 0; fi19 < 13; fi19++ {
            switch fi19 {
            case 1:
                if ((v2 % ((1) | 1)) > v2) {
                    delete(v1, 2147483648)
                    v2 -= (^(v2))
                }
            case 9223372036854775807:
                var v20 uint16 = 42
                _ = v20
            case -9223372036854775808:
                v2 = v2
                var v21 bool = (((!false) && false) || (!(!("foo bar" > "hello"))))
                _ = v21
            }
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    v2 = v2
    for fi22 := 0; fi22 < 3; fi22++ {
        fmt.Println(fi22)
        v23 := v1[42]
        _ = fi22
        _ = v23
    }
    if (!((815 >= 42) && true)) {
        if (!true) {
            v2 += (v2 >> ((v2) & 63))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

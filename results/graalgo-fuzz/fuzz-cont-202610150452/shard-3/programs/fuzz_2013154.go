package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int16; F2 uint32 }

func main() {
    var v1 uint32 = 65535
    v2 := []int{4294967295}
    v3 := make(map[uint32]string)
    var v4 int64 = 178
    v5 := GenS1{F0: (v4 + v4)}
    v4 = 42
    v4 += (v4 << ((65535) & 63))
    v1 += v1
    _ = len(v3)
    v1 += (v1 | v1)
    var v6 int16 = -1
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS2{F0: true}
            var v10 bool = true
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    var v11 uint = 4294967295
    v12 := GenS3{F0: 42, F1: (v6 >> (((v11 + 16)) & 63)), F2: v1}
    for fi13 := 0; fi13 < 18; fi13++ {
        switch "foo bar" {
        case "":
            var v14 int = fi13
            _ = v14
        case "a":
            if false {
                var v15 int64 = (v4 >> (((^(v11))) & 63))
                if true {
                    _ = len(v3)
                    var v16 int64 = (v4 % (((v4 ^ v4)) | 1))
                    v17 := GenS3{F0: v1, F1: v6, F2: v1}
                    _ = v16
                    _ = v17
                } else {
                    var v18 bool = ((0.1 + 1e10) == -2.25)
                    _ = v18
                }
                fmt.Println(fi13, v11)
                _ = v15
            } else {
                if true {
                    var v19 float32 = (3.14159 + 0.1)
                    _ = v19
                }
            }
        case "\n\t":
            for fi20 := 0; fi20 < 18; fi20++ {
                v6 = (v6 | (^(v6)))
                _ = fi20
            }
        }
        switch ("foo bar" + ("hello" + "a")) {
        case "":
            v6 = -1
        case "foo bar":
            if (false && ("\n\t" <= ((("a" + "a") + "") + ""))) {
                v21 := GenS2{F0: (105 != -128)}
                _ = v21
            } else {
                fmt.Println(v6, fi13)
                if (!((^(v1)) == 542)) {
                    switch 32767 {
                    case 9223372036854775807:
                        v1 += (v1 << (((v11 % (((v11 ^ 255)) | 1))) & 63))
                    }
                    if false {
                        v22 := GenS3{F0: (v1 ^ v1), F1: (v6 % (((v6 + v6)) | 1)), F2: v1}
                        _ = v22
                    } else {
                        _ = len(v3)
                    }
                }
            }
        }
        v23 := GenS2{F0: (int(v6) >= -151)}
        _ = fi13
        _ = v23
    }
    v11 -= 255
    v24 := GenS2{F0: true}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v11)
    var v25 int8 = 41
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

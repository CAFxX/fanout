package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }

func main() {
    var v1 bool = false
    var v2 float32 = 1e10
    var v3 int32 = 42
    var v4 string = "foo bar"
    var v5 int8 = 1
    v3 = v3
    var v6 float32 = v2
    var v7 int64 = -9223372036854775808
    if true {
        fmt.Println(v6, v1, v1)
    }
    var v8 int8 = v5
    if true {
        if (v1 || v1) {
            var v9 float32 = v2
            v10 := GenS1{F0: "\n\t"}
            _ = v10
            _ = v9
        } else {
            if (4294967295 != 4294967295) {
                v11 := []int{-1, 127, 452, -791}
                v12 := GenS2{F0: v4}
                _ = v11
                _ = v12
            }
            switch v4 {
            case "hello":
                switch (v4 + v4) {
                case "":
                    var v13 int32 = 42
                    _ = v13
                default:
                    if false {
                        var v14 int16 = 32767
                        for fi15 := 0; fi15 < 10; fi15++ {
                            v16 := []int{(fi15 % (((fi15 | 9223372036854775807)) | 1)), -1}
                            if true {
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                            }
                            if v1 {
                                for fi17 := 0; fi17 < 1; fi17++ {
                                    v18 := GenS3{F0: (4294967295 <= 283)}
                                    _ = fi17
                                    _ = v18
                                }
                            }
                            _ = fi15
                            _ = v16
                        }
                        _ = v14
                    }
                }
            }
        }
        v4 = v4
    }
    v3 = (v3 | (v3 >> ((1) & 63)))
    var v19 float32 = (v6 / v6)
    fmt.Println(v7)
    fmt.Println(v6)
    v5 += (v8 - (v5 + 127))
    v20 := GenS3{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}

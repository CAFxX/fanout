package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
func genFn0() uint {
    var v5 uint = 2
    switch (("foo bar" + "foo bar") + ("" + "")) {
    case "hello":
        v5 = v5
    case "a":
        if (1 >= 1) {
            switch 9223372036854775807 {
            case 127:
                switch 521 {
                case 4294967295:
                    v6 := GenS1{F0: 0}
                    _ = v6
                case -567:
                    fmt.Println(v5)
                }
            }
        }
    case "":
        if (!(float32(0.1) < (3.14159 - (0.0 / (-2.25 * (-2.25 - 1e10)))))) {
            v7 := []int{32767, -1}
            v5 = 4294967295
            _ = v7
        } else {
            switch "hello" {
            case "hello":
                for fi8 := 0; fi8 < 10; fi8++ {
                    if true {
                        var v9 uint8 = 255
                        _ = v9
                    }
                    v5 = (v5 / (((v5 / ((v5) | 1))) | 1))
                    _ = fi8
                }
            }
            var v10 uint = v5
            _ = v10
        }
    default:
        var v11 int64 = 0
        _ = v11
    }
    var v12 bool = true
    var v13 uint8 = 127
    _ = v12
    _ = v13
    _ = v5
    return v5
}

func main() {
    v1 := []int{-539, 76, -317, 1}
    var v2 string = ""
    v3 := make(map[string]string)
    var v4 int = 2147483647
    v2 = v2
    v4 = v4
    v14 := GenS1{F0: 0}
    fmt.Println(v4, v4)
    if (!true) {
        v15 := GenS2{F0: 0}
        switch v2 {
        case "":
            v16 := GenS1{F0: -1}
            _ = v16
        case "hello":
            v17 := GenS2{F0: 4294967295}
            _ = v17
        }
        _ = v15
    }
    if true {
        if (-9223372036854775808 != -207) {
            v4 -= 2147483647
            _ = len(v1)
            v3[(((v2 + (v2 + "foo bar")) + v2) + "a")] = ("\n\t" + (v2 + (((v2 + "a") + v2) + (v2 + v2))))
        } else {
            var v18 uint16 = 255
            v4 = (v4 / (((v4 * v4)) | 1))
            for fi19 := 0; fi19 < 19; fi19++ {
                if true {
                    v20 := v3[(v2 + "")]
                    var v21 uint64 = 939
                    _ = v20
                    _ = v21
                }
                switch (("" + "") + "a") {
                case "a":
                    fmt.Println(v4, fi19, v2)
                case "":
                    var v22 uint16 = (v18 / (((v18 & v18)) | 1))
                    _ = v22
                }
                _ = fi19
            }
            _ = v18
        }
    }
    var v23 string = v2
    v24 := GenS2{F0: genFn0()}
    v4 = (v4 << ((genFn0()) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 109
    v3 := make(map[string]bool)
    v3["foo bar"] = (false || true)
    if false {
        if true {
            _ = len(v3)
            for fi4 := 0; fi4 < 8; fi4++ {
                for fi5 := 0; fi5 < 6; fi5++ {
                    var v6 bool = (!false)
                    _ = fi5
                    _ = v6
                }
                v1["\n\t"] = fi4
                _ = fi4
            }
        }
    } else {
        v2 = v2
        switch "\n\t" {
        case "hello":
            if false {
                var v7 uint16 = 1
                v8 := []int{32768, -1, 2147483647}
                _ = v7
                _ = v8
            }
        default:
            v2 += v2
        }
    }
    var v9 int = 42
    v10 := []int{128}
    delete(v3, ("hello" + "a"))
    v9 += (-(v9))
    v9 -= v9
    v9 = (v9 * (v9 * v9))
    var v11 int64 = 0
    var v12 int = (v9 ^ (-(v9)))
    v2 = v2
    v13 := v1["hello"]
    var v14 float32 = 0.1
    switch 2147483647 {
    case 32768:
        for fi15 := 0; fi15 < 20; fi15++ {
            for fi16 := 0; fi16 < 5; fi16++ {
                if false {
                    fmt.Println(v14)
                } else {
                    v11 = (v11 * v11)
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                }
                var v17 string = "a"
                _ = fi16
                _ = v17
            }
            if true {
                v18 := GenS1{F0: 480}
                v9 -= v13
                var v19 int32 = -345
                _ = v18
                _ = v19
            } else {
                delete(v3, "hello")
                var v20 uint8 = (^(v2))
                _ = v20
            }
            for fi21 := 0; fi21 < 19; fi21++ {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                v22 := v3["foo bar"]
                v23 := GenS1{F0: 18446744073709551615}
                v12 = v9
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi15
        }
    case 0:
        if (false && (!(2 >= 982))) {
            v14 += v14
            v11 = (v11 << ((4294967295) & 63))
        }
    case 4294967295:
        fmt.Println(v9, v14)
    default:
        switch (v12 % (((v9 - v9)) | 1)) {
        case 345:
            v13 = int(v2)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []string{"a"}
    v2 := []string{}
    var v3 uint64 = 4294967295
    var v4 uint16 = 32767
    var v5 string = "a"
    var v6 string = v5
    _ = len(v1)
    if false {
        v2 = append(v2, "hello")
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v7 uint64 = 1
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            if (127 < 36) {
                var v10 uint16 = v4
                v2 = append(v2, ((("\n\t" + (v6 + (v6 + ((("" + "\n\t") + ("hello" + v6)) + (v6 + ((v6 + "\n\t") + "hello")))))) + "hello") + (("" + ("\n\t" + v6)) + ("\n\t" + v6))))
                _ = v10
            } else {
                switch (fi8 / (((fi8 * -672)) | 1)) {
                case 42:
                    var v11 int16 = 365
                    _ = v11
                default:
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(v3, v7)
                    }
                }
                var v12 int16 = 0
                _ = v12
            }
            _ = fi9
        }
        _ = fi8
    }
    _ = len(v1)
    v13 := GenS1{F0: -48}
    v3 -= 507
    fmt.Println(v5, v4)
    v5 = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
